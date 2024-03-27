BEGIN;

CREATE TYPE status_cpf_validation_enum AS ENUM(
    'PENDING',
    'REPROVED',
    'APPROVED',
    'INVALID_BIRTHDAY'
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT false,
    cpf VARCHAR(30) NOT NULL UNIQUE,
    birthday DATE NOT NULL,
    status_cpf_validation status_cpf_validation_enum NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "user" IS 'Table of user the system';

COMMENT ON COLUMN "user".id IS 'Primary key of table';

COMMENT ON COLUMN "user".name IS 'Name of user';

COMMENT ON COLUMN "user".email IS 'Name of user';

COMMENT ON COLUMN "user".password IS 'Hash to password of user';

COMMENT ON COLUMN "user".active IS 'Indicates if the is active';

COMMENT ON COLUMN "user".cpf IS 'Document of user';

COMMENT ON COLUMN "user".birthday IS 'birthday of user';

COMMENT ON COLUMN "user".status_cpf_validation IS 'Status of validation to cpf of user';

COMMIT;