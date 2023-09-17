package sources

import (
	"MidhunRajeevan/payments/config"
	"context"
	"log"
)

func execAndScan(statement string, params []interface{}) (SourceSystem, error) {
	var w SourceSystem
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.Did, &w.StandardName, &w.NameEn, &w.NameAr, &w.Description, &w.Status, &w.IsActive,
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

func execAndScanToList(statement string, params []interface{}) ([]SourceSystem, error) {
	w := make([]SourceSystem, 0)
	rows, err := config.DB.Query(context.Background(), statement, params...)
	if err != nil {
		log.Println("Database Select Error:", err.Error())
	}
	for rows.Next() {
		r := SourceSystem{}
		err = rows.Scan(&r.ID, &r.Did,
			&r.StandardName, &r.NameEn, &r.NameAr, &r.Description, &r.Status,
			&r.IsActive, &r.IsArchived, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		w = append(w, r)
	}
	return w, err
}

func SelectSource(did string) (SourceSystem, error) {
	statement := `
	select
  	id, did, standardName, name_en, name_ar, description, status, is_active,
	is_archived, created_at, updated_at
	from source_systems
	where did = $1`
	params := make([]interface{}, 0)
	params = append(params, did)
	return execAndScan(statement, params)
}

func SelectSources() ([]SourceSystem, error) {
	statement := `
	select id, did, standardName, name_en, name_ar, description, status, is_active,
	is_archived, created_at, updated_at
	from source_systems
	where is_active =true`
	params := make([]interface{}, 0)
	return execAndScanToList(statement, params)
}

func selectSourceByCode(did string) (SourceSystem, error) {
	statement := `
	select id, did, standardName, name_en, name_ar, description, status, is_active,
	is_archived, created_at, updated_at
	from source_systems
	where did = $1`
	params := make([]interface{}, 0)
	params = append(params, did)
	return execAndScan(statement, params)
}

func upsertSource(r SourceSystem) (SourceSystem, error) {
	statement := `
	insert into source_systems
		(did, standardname, name_en, name_ar, description)
	values ($1, $2, $3, $4, $5)
	on conflict (did) do update
		set standardname = $2,
		name_en = $3,
		name_ar = $4,
		description = $5,
		updated_at = current_timestamp
		where source_systems.did = $1
	returning
		id, did, standardname, name_en, name_ar, description, 
		status, is_active, is_archived, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.Did, r.StandardName, r.NameEn, r.NameAr, r.Description)
	return execAndScan(statement, params)
}

func updateSourceStatus(r SourceSystem) (SourceSystem, error) {
	statement := `
	update source_systems set
  	is_active = $1,
	updated_at = current_timestamp
	where did = $2
	returning
		id, did, standardname, name_en, name_ar, description, 
		status, is_active, is_archived, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params, r.IsActive, r.Did)
	return execAndScan(statement, params)
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
