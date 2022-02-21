package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant/database"
	"restaurant/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while lisiting invoice items"})

		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Params("invoice_id")
		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while lisitng items"})
		}
		var invoiceView InvoiceViewFormat
		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date
		invoiceView.payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.payment_method = *&invoice.Payment_method
		}
		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Payment_due = allOrderItems[0]["table_number"]
		invoiceView.Order_details = allOrderItems[0]["order_items"]
		c.JSON(http.StatusOK, invoiceVie)
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel=context.WithTimeout(context.Background(),100*time.Second)
		var invoice models.Invoice
		if err:=c.BindJSON(&invoice);err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})		
			return	
		}
		var order models.Order
		err:=orderCollection.FindOne(ctx,bson.M{"order_id":invoice.Invoice_id}).Decode(&order)
		defer cancel()
		if err != nil {
			msg.Sprintf("message: Order was not found")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice
		invoiceId := c.Params("invoice_id")
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		filter := bson.M{"invoice_id": invoiceId}
		var UpdateObj primitive.D
		if invoice.Payment_method != nil {

		}
		if invoice.Payment_status != nil {

		}
		invoice.Updated_at,_=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		UpdateObj = append(UpdateObj, bson.E{"updated_at",invoice.Updated_at})
		upsert:=true
		opt:=options.UpdateOptions{
			Upsert: &upsert,
		}
		status:="PENDING"
		if invoice.Payment_status==nil{
			invoice.Payment_status=&status
		}
		result,err:= invoiceCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set",UpdateObj},
			},
			&opt,
		)
		if err!=nil{
			msg:=fmt.Sprintf("invoice item update failed")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
}