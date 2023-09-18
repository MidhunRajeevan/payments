package config

import (
	"context"
	"log"
)

var statements []string

func build() {
	statements = make([]string, 0)

	/*
	 * citext extension
	 */
	statements = append(statements, `
		create extension if not exists citext with schema public`)

	/*
	 table for payment requests
	*/
	statements = append(statements, `
		create table if not exists requests (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		payload jsonb not null default '[]'::json,
		token text,
		transactionReference text,
		status  text not null default 'created',
		attempts    int not null default 0,
		message     text,
		is_expired boolean NOT NULL DEFAULT false,
		payment_initiated_at      timestamptz,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)`)

	/*
	 table for source systems , should have name and did , created and updated dates
	 source system can be activate or deactivate , there should be a audit table to track the same
	*/
	statements = append(statements, `
		create table if not exists source_systems (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		standardName citext NOT NULL,
		name_en citext NOT NULL,
		name_ar citext NOT NULL,
		description text,
		status  text not null default 'active',
		is_active boolean NOT NULL DEFAULT true,
		is_archived boolean not null default false,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)`)

	statements = append(statements, `
		create unique index if not exists ux_did_on_source_systems
		on source_systems
		using btree (did)`)
	/*
	 table to capture the configuration details of the source system
	*/
	statements = append(statements, `
		create table if not exists source_system_config (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		sp_code text NOT NULL,
		sp_value text,
		rms_source_id text,
		rms_source_name text,
		rms_source_subfield text,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)`)

	/*
	 static table for payment gatways
	 should be able to enable disable gateway for the entire solution
	*/
	statements = append(statements, `
		create table if not exists gateways (
		id serial PRIMARY KEY,
		code citext NOT NULL,
		standardName citext NOT NULL,
		name_en citext NOT NULL,
		name_ar citext NOT NULL,
		description text,
		status  text not null default 'active',
		is_active boolean NOT NULL DEFAULT true,
		is_archived boolean not null default false,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)
	`)
	statements = append(statements, `
		create unique index if not exists ux_code_on_gateways
		on gateways
		using btree (code)`)

	/*
	 source_system_payment_gateway table maps the source systme with gateway
	*/
	statements = append(statements, `
		create table if not exists source_system_gateway (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		gateway_code citext NOT NULL,
		status  text not null default 'active',
		is_active boolean NOT NULL DEFAULT true,
		is_archived boolean not null default false,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)
	`)

	statements = append(statements, `
		create unique index if not exists idx_unique_did_gateway_code
		on source_system_gateway
		using btree (did, gateway_code)`)

	/*
	 source_system_audit table tracks the changes to source_system,source_system_config,source_system_payment_gateway
	*/
	statements = append(statements, `
		create table if not exists source_system_audit (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		data jsonb not null default '{}'::json,
		created_at timestamptz NOT NULL DEFAULT now()
	)
	`)

	/*
	 customer to capture the customer associated with the invoice
	*/
	statements = append(statements, `
		create table if not exists customer (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		name text,
		custgroup text,
		email citext,
		trade_license_number text NOT NULL, 
		account text,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)
	`)

	/*
	 invoice table to capture the invoice details of the source system
	*/
	statements = append(statements,
		`create table if not exists invoice (
		id  bigserial PRIMARY KEY,
		number      citext NOT NULL,
		issued_date date NOT NULL,
		last_date   date NOT NULL,
		did citext NOT NULL,
		customer_id int,
		amount_value float NOT NULL,
		amount_currency citext NOT NULL,
		payment_status citext NOT NULL default 'not-paid',
		payment_date timestamptz,
		purchase_order_form_no text,
		document_date timestamptz,
		rms_invoice_date timestamptz,
		company_traffic_file_number citext NOT NULL,
		company_trade_license_number citext NOT NULL,
		created_at  timestamptz NOT NULL DEFAULT now(),
		updated_at  timestamptz NOT NULL DEFAULT now()
	)`)

	statements = append(statements, `
		create table if not exists transactions (
		id bigserial PRIMARY KEY,
		did citext NOT NULL,
		invoice_id int NOT NULL, 
		gateway_id int NOT NULL,
		reference_number citext NOT NULL,
		request_payload jsonb not null default '[]'::json,
		response_payload jsonb not null default '[]'::json,
		amount numeric NOT NULL,
		payment_requested_at timestamptz NOT NULL DEFAULT now(),
		payment_confirmation_at timestamptz,
		payment_status citext NOT NULL default 'requested',
		payment_message text,
		created_at timestamptz NOT NULL DEFAULT now(),
		updated_at timestamptz NOT NULL DEFAULT now()
	)
	`)

}

// Setup database
func Setup() {
	build()
	log.Println("Creating database objects...")
	for _, statement := range statements {
		if _, err := DB.Exec(context.Background(), statement); err != nil {
			log.Println("Database query failed!")
			log.Fatalln(statement)
		}
	}
}
