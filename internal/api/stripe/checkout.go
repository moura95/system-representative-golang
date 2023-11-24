package stripe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"

	"my-orders/internal/util"
)

func checkout(email string, name string, ProductID string, RepresentativeID int32, coupon string) (*stripe.CheckoutSession, error) {
	var discounts []*stripe.CheckoutSessionDiscountParams
	if coupon == "MIDAS20" {
		discounts = append(discounts, &stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String("thp0lt0r"),
		})
	}
	if coupon == "MIDAS20F" {
		discounts = append(discounts, &stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String("KNlQO6Qe"),
		})
	}

	representativeString := fmt.Sprintf("%d", RepresentativeID)
	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("Email", email)
	customerParams.AddMetadata("CustomerName", name)
	customerParams.AddMetadata("RepresentativeID", representativeString)
	customerParams.AddMetadata("System", "MIDASREP")
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, err
	}

	meta := map[string]string{
		"Email":            email,
		"RepresentativeID": representativeString,
		"CustomerName":     name,
	}

	log.Println("Creating meta for user: ", meta)

	params := &stripe.CheckoutSessionParams{
		Customer:   &newCustomer.ID,
		SuccessURL: stripe.String("https://midasgestor.com.br/checkout/success"),
		CancelURL:  stripe.String("https://midasgestor.com.br/checkout/failed"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Discounts: discounts,
		Mode:      stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(ProductID),
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Metadata: meta,
		},
	}
	return session.New(params)
}

type checkoutInput struct {
	Plan   string `form:"plan"`
	Coupon string `form:"coupon"`
}

type SessionOutput struct {
	Id string `json:"id"`
}

func (s *Stripe) checkoutCreator(ctx *gin.Context) {
	input := &checkoutInput{}
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	var productID string
	representativeID := ctx.Keys["representativeID"].(int32)
	// get email ctx
	user, err := s.Db.GetEmailAndNameByRepresentativeID(ctx, representativeID)
	email := user.Email.String
	name := user.Name.String
	if err != nil {
		fmt.Println(err)
	}
	if input.Plan == "silver" {
		productID = s.Config.SilverPlanID
	} else {
		productID = s.Config.GoldPlanID
	}
	stripeSession, err := checkout(email, name, productID, 1, input.Coupon)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewEncoder(ctx.Writer).Encode(&SessionOutput{Id: stripeSession.ID})

	if err != nil {
		fmt.Println(err)
	}
}
