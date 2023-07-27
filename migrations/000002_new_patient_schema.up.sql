CREATE TABLE IF NOT EXISTS "patient" (
         "id" varchar NOT NULL,
         "name" varchar NOT NULL,
         "gender" varchar NOT NULL,
         "phone" varchar NOT NULL,
         "email" varchar  NOT NULL,
         "age" INT NOT NULL,
         "city" varchar NOT NULL,
         "state" varchar NOT NULL,
         "pincode" varchar NOT NULL,
         "joined_on" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
         "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
         "updated_at" TIMESTAMP WITHOUT TIME ZONE,
         "last_login_time" TIMESTAMP WITHOUT TIME ZONE,
         "meta_data" varchar,
         "status" varchar  NOT NULL,
         "password" varchar NOT NULL,
         "salt" varchar NOT NULL,
         PRIMARY KEY ("id")
)