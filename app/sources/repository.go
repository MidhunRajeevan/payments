package sources

import (
	"MidhunRajeevan/payments/config"
	"context"
	"log"
)

func execAndScanSourceSystem(statement string, params []interface{}) (SourceSystem, error) {
	var w SourceSystem
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.Did, &w.NameEn, &w.NameAr, &w.Status, &w.IsActive,
		&w.IsArchived, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func execAndScanSourceSystemGateway(statement string, params []interface{}) (SourceSystemGateway, error) {
	var w SourceSystemGateway
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.DID, &w.GatewayCode, &w.Status, &w.IsActive,
		&w.IsArchived, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func SelectSource(did string) (SourceSystem, error) {
	statement := `
	select
  	id, did, name_en, name_ar, status, is_active,
	is_archived, created_at, updated_at
	from source_system
	where did = $1 and is_active = true`
	params := make([]interface{}, 0)
	params = append(params, did)
	return execAndScanSourceSystem(statement, params)
}

func insertSourceSystemGateway(r SourceSystemGateway) (SourceSystemGateway, error) {
	statement := `
	insert into source_system_gateway
	(did, gateway_code)
	values ($1, $2)
	returning
		id, did, gateway_code,status, is_active, is_archived, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.DID, r.GatewayCode)
	return execAndScanSourceSystemGateway(statement, params)
}
