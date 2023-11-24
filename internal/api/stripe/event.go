package stripe

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/stripe/stripe-go/v72/webhook"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

func (s *Stripe) handlerEvent(ctx *gin.Context) {
	const MaxBodyBytes = int64(65536)
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	endpointSecret := s.Config.StripeWebhookSecret
	event, err := webhook.ConstructEvent(payload, ctx.Request.Header.Get("Stripe-Signature"),
		endpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		ctx.JSON(http.StatusOK, err)
		return
	}
	fmt.Printf("event: %s \n", event.Type)

	switch event.Type {
	case "invoice.paid":
		var invoicePaid stripe.Invoice
		err := json.Unmarshal(event.Data.Raw, &invoicePaid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			ctx.JSON(http.StatusOK, err)
			return
		}

		stripeID := event.ID
		representativeID := invoicePaid.Lines.Data[0].Metadata["RepresentativeID"]
		representativeIDInt, err := strconv.ParseInt(representativeID, 10, 32)
		intervalCount := invoicePaid.Lines.Data[0].Plan.IntervalCount
		dateExpire := time.Now().AddDate(0, int(intervalCount), 0)

		arg := repository.UpdatePlanByIDParams{
			ID:   int32(representativeIDInt),
			Plan: repository.PlanTypesSilver,
			StripeID: sql.NullString{
				String: stripeID,
				Valid:  true,
			},
			DataExpire: dateExpire,
		}

		_, err = s.Db.UpdatePlanByID(ctx, arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			ctx.JSON(http.StatusOK, err)
			return
		}
		fmt.Printf("Paid was successful! UserID: %s \n", representativeID)
		email := invoicePaid.Lines.Data[0].Metadata["Email"]
		name := invoicePaid.Lines.Data[0].Metadata["CustomerName"]
		data := util.Data{
			Title: "Pagamento Efetuado Com Sucesso",
			Name:  name,
			Msg:   "Seu Pagamento foi efetuado com sucesso, agora você tem acesso ao plano ",
		}

		err = util.NewSender("MidasGestor", repository.Smtp{}).SendEmail("Pagamento com sucesso", data, "payment-success",
			[]string{email}, nil, nil, nil,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending email: %v\n", err)
			util.ErrorResponse(400, "", "Error ao enviar email")
			return
		}
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
		return
	case "invoice.payment_failed":
		var invoicePaid stripe.Invoice
		err := json.Unmarshal(event.Data.Raw, &invoicePaid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			ctx.JSON(http.StatusOK, err)
			return
		}
		representativeID := invoicePaid.Lines.Data[0].Metadata["RepresentativeID"]
		fmt.Printf("Payment Failled! to UserID: %s \n", representativeID)
		email := invoicePaid.Lines.Data[0].Metadata["Email"]
		name := invoicePaid.Lines.Data[0].Metadata["CustomerName"]
		data := util.Data{
			Title: "Pagamento Falhou",
			Name:  name,
			Msg:   "Seu Pagamento foi recusado tente novamente pelo sistema.",
		}

		err = util.NewSender(
			"MidasGestor", repository.Smtp{}).SendEmail("Pagamento Falhou", data, "payment-failed",
			[]string{email}, nil, nil, nil,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending email: %v\n", err)
			util.ErrorResponse(400, "", "Error ao enviar email")
			return
		}
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
		return
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
