CREATE TABLE public."ORDER_ITEM" (
    "ORDER_ID" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "PRODUCT_ID" UUID NOT NULL,
    "TRANSACTION_ID" UUID NOT NULL,
    "USER_ID" UUID NOT NULL,
    "DELIVERY_STATUS" VARCHAR(60) NOT NULL,
    "PAYMENT_STATUS" VARCHAR(60) NOT null,
    "CREATED_AT" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "UPDATED_AT" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE "TRANSACTION" (
    "TRANSACTION_ID" UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Assuming TRANSACTION_ID is unique
    "STATUS" VARCHAR(50) NOT NULL, -- To ensure only valid status values
    "PROCESSING" INT NOT NULL, -- Integer to store processing status
    "CREATED_AT" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Auto-populate with current timestamp
    "UPDATED_AT" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "OUTBOX" (
    "AGGREGATE_ID" uuid NOT NULL,   -- Corresponds to AggregateID
    "COMMAND_ID" uuid NOT NULL,    -- Corresponds to CommandID
    "COMMAND_TYPE" VARCHAR(255) NOT NULL,  -- Corresponds to CommandType
    "PAYLOAD" text,                       -- Corresponds to Payloay, stored as JSON for flexibility
    "REPLY_TO" VARCHAR(255),               -- Corresponds to ReplyTo
    PRIMARY KEY ("COMMAND_ID")             -- Unique identifier for each command
);
