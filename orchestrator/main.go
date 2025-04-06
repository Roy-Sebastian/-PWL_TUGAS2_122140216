package main

import (
	"context"
	"time"

	orderpb "order-service/proto"
	paymentpb "payment-service/proto"
	shippingpb "shipping-service/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// Configure logrus
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to Order Service
	orderConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to order service")
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// Connect to Payment Service
	paymentConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to payment service")
	}
	defer paymentConn.Close()
	paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)

	// Connect to Shipping Service
	shippingConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to shipping service")
	}
	defer shippingConn.Close()
	shippingClient := shippingpb.NewShippingServiceClient(shippingConn)

	// START SAGA
	orderID := "12345"
	log.WithField("order_id", orderID).Info("Starting SAGA transaction")

	// 1. Create Order
	_, err = orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{OrderId: orderID})
	if err != nil {
		log.WithError(err).Fatal("Failed to create order")
	}
	log.WithField("order_id", orderID).Info("Order created successfully")

	// 2. Process Payment
	_, err = paymentClient.ProcessPayment(ctx, &paymentpb.ProcessPaymentRequest{OrderId: orderID})
	if err != nil {
		log.WithError(err).Warn("Payment failed. Starting compensation...")
		_, cancelErr := orderClient.CancelOrder(ctx, &orderpb.CancelOrderRequest{OrderId: orderID})
		if cancelErr != nil {
			log.WithError(cancelErr).Error("Failed to cancel order")
		}
		return
	}
	log.WithField("order_id", orderID).Info("Payment processed successfully")

	// 3. Start Shipping
	_, err = shippingClient.StartShipping(ctx, &shippingpb.StartShippingRequest{OrderId: orderID})
	if err != nil {
		log.WithError(err).Warn("Shipping failed. Starting compensation...")

		// Cancel Shipping
		_, cancelErr := shippingClient.CancelShipping(ctx, &shippingpb.CancelShippingRequest{OrderId: orderID})
		if cancelErr != nil {
			log.WithFields(logrus.Fields{
				"order_id": orderID,
				"error":    cancelErr,
			}).Error("Failed to cancel shipping")
		} else {
			log.WithField("order_id", orderID).Info("Shipping cancelled successfully")
		}

		// Refund Payment
		_, refundErr := paymentClient.RefundPayment(ctx, &paymentpb.RefundPaymentRequest{OrderId: orderID})
		if refundErr != nil {
			log.WithFields(logrus.Fields{
				"order_id": orderID,
				"error":    refundErr,
			}).Error("Failed to refund payment")
		} else {
			log.WithField("order_id", orderID).Info("Payment refunded successfully")
		}

		// Cancel Order
		_, cancelOrderErr := orderClient.CancelOrder(ctx, &orderpb.CancelOrderRequest{OrderId: orderID})
		if cancelOrderErr != nil {
			log.WithFields(logrus.Fields{
				"order_id": orderID,
				"error":    cancelOrderErr,
			}).Error("Failed to cancel order")
		} else {
			log.WithField("order_id", orderID).Info("Order cancelled successfully")
		}

		return
	}
	log.WithField("order_id", orderID).Info("Shipping started successfully")

	log.WithField("order_id", orderID).Info("SAGA completed successfully")
}
