package payment

import (
	"MidhunRajeevan/payments/app/util"
	"log"
	"net/http"

	pgx "github.com/jackc/pgx/v4"
)

func getCompanies(w http.ResponseWriter, req InvoiceRequest) (Companies, bool) {

	var company Companies
	var err error
	var ok bool

	company.Name.En = req.Company.Name.En
	company.Name.Ar = req.Company.Name.Ar
	company.TradeLicenseNumber = req.Company.TradeLicenseNumber
	company.TrafficFileNumber = req.Company.TrafficFileNumber
	for _, email := range req.Company.Contact.Emails {
		if email.Preferred && len(email.Email) != 0 {
			company.Email = email.Email
			break
		}
	}
	dbcompany, err := selectCompanies(req.Company.TradeLicenseNumber)
	if err == pgx.ErrNoRows {
		company, ok = createCompanies(w, company)
		if !ok {
			return company, false
		}

	} else if err != nil {

		log.Println("Select Companies Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return company, false

	} else if company.Email != dbcompany.Email {
		company, err = upsertCompaniesTable(company)
		if err != nil {
			log.Println("Insert to customer table Error:", err.Error())
			util.InternalServerError(&w, "contact_support")
			return company, false
		}
	} else {
		company = dbcompany
	}
	return company, true
}

// if the customer exist in RMS , then get RMS customer and upsert to DB
// if customer does not exist in RMS , then create new customer and upsert to DB
func createCompanies(w http.ResponseWriter, reqCompanies Companies) (Companies, bool) {

	var company Companies
	var err error
	// check if the customer exist in RMS
	rmscustomer, isExist := getCompaniesRMS(reqCompanies.TradeLicenseNumber)
	if isExist {
		reqCompanies.Account = rmscustomer.Account
		reqCompanies.CustGroup = rmscustomer.CustGroup
		company, err = upsertCompaniesTable(reqCompanies)
		if err != nil {
			log.Println("Insert to customer table Error:", err.Error())
			util.InternalServerError(&w, "contact_support")
			return company, false
		}
	} else {
		// throw exception if the get RMS customer failed
		if len(rmscustomer.Account) == 0 {
			log.Println("Get RMS Companies Error")
			util.InternalServerError(&w, "contact_support")
			return company, false
		} else {
			// create customer in RMS if not exist
			rmsCreatedCust, iscreate := createCompaniesRMS(company)
			if !iscreate {
				log.Println("RMS customer creation Error")
				util.InternalServerError(&w, "contact_support")
				return company, false
			} else {
				company.Account = rmsCreatedCust.Account
				company.CustGroup = rmsCreatedCust.CustGroup
				company, err = upsertCompaniesTable(company)
				if err != nil {
					log.Println("Insert to customer table Error:", err.Error())
					util.InternalServerError(&w, "contact_support")
					return company, false
				}
			}
		}
	}
	return company, true

}

func getCompaniesRMS(tradeLicenseNumber string) (Companies, bool) {

	var customer Companies
	customer.CustGroup = "group"
	customer.Account = "122222"
	return customer, true
}

func createCompaniesRMS(customer Companies) (Companies, bool) {

	customer.CustGroup = "group"
	customer.Account = "122222"
	return customer, true
}

func generateSalesOrderRMS(invoice InvoiceRequest, company Companies) (string, bool) {
	return "1425262", true
}
