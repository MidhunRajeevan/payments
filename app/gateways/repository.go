package gateways

import (
	"MidhunRajeevan/payments/config"
	"context"
	"log"
)

func execAndScanToList(statement string, params []interface{}) ([]PaymentGateway, error) {
	w := make([]PaymentGateway, 0)
	rows, err := config.DB.Query(context.Background(), statement, params...)
	if err != nil {
		log.Println("Database Select Error:", err.Error())
	}
	for rows.Next() {
		r := PaymentGateway{}
		err = rows.Scan(&r.ID, &r.Code,
			&r.StandardName, &r.NameEn, &r.NameAr, &r.Description, &r.Status,
			&r.IsActive, &r.IsArchived, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		w = append(w, r)
	}
	return w, err
}

func execAndScan(statement string, params []interface{}) (PaymentGateway, error) {
	var w PaymentGateway
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.Code, &w.StandardName, &w.NameEn, &w.NameAr, &w.Description, &w.Status, &w.IsActive,
		&w.IsArchived, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func SelectGatewaysByDid(did string) ([]PaymentGateway, error) {
	statement := `
	select g.id,g.code, g.standardName, g.name_en, g.name_ar, g.description, 
	g.status, g.is_active, g.is_archived,g.created_at, g.updated_at
	from gateways g
	inner join source_system_gateway sg
		  on g.code = sg.gateway_code
	where sg.did = $1 and g.is_active =true`
	params := make([]interface{}, 0)
	params = append(params, did)
	return execAndScanToList(statement, params)
}

func SelectGateways() ([]PaymentGateway, error) {
	statement := `
	select g.id,g.code, g.standardName, g.name_en, g.name_ar, g.description, 
	g.status, g.is_active, g.is_archived,g.created_at, g.updated_at
	from gateways g
	where g.is_active =true`
	params := make([]interface{}, 0)
	return execAndScanToList(statement, params)
}

func SelectGatewayByCode(gatewayCode string) (PaymentGateway, error) {
	statement := `
	select g.id,g.code, g.standardName, g.name_en, g.name_ar, g.description, 
	g.status, g.is_active, g.is_archived,g.created_at, g.updated_at
	from gateways g
	where g.code = $1`
	params := make([]interface{}, 0)
	params = append(params, gatewayCode)
	return execAndScan(statement, params)
}

func upsertGateway(r PaymentGateway) (PaymentGateway, error) {
	statement := `
	insert into gateways
		(code, standardname, name_en, name_ar, description)
	values ($1, $2, $3, $4, $5)
	on conflict (code) do update
		set standardname = $2,
		name_en = $3,
		name_ar = $4,
		description = $5,
		updated_at = current_timestamp
		where gateways.code = $1
	returning
		id, code, standardname, name_en, name_ar, description, 
		status, is_active, is_archived, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.Code, r.StandardName, r.NameEn, r.NameAr, r.Description)
	return execAndScan(statement, params)
}

func updateGatewayStatus(r PaymentGateway) (PaymentGateway, error) {
	statement := `
	update gateways set
  	is_active = $1,
	updated_at = current_timestamp
	where code = $2
	returning
		id, code, standardname, name_en, name_ar, description, 
		status, is_active, is_archived, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params, r.IsActive, r.Code)
	return execAndScan(statement, params)
}
