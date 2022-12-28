CREATE TABLE
  public.customers (
    customerid serial NOT NULL,
    firstname character varying(50) NOT NULL,
    lastname character varying(50) NOT NULL,
    address1 character varying(50) NOT NULL,
    address2 character varying(50) NULL,
    city character varying(50) NOT NULL,
    state character varying(50) NULL,
    zip integer NULL,
    country character varying(50) NOT NULL,
    region smallint NOT NULL,
    email character varying(50) NULL,
    phone character varying(50) NULL,
    creditcardtype integer NOT NULL,
    creditcard character varying(50) NOT NULL,
    creditcardexpiration character varying(50) NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(50) NOT NULL,
    age smallint NULL,
    income integer NULL,
    gender character varying(1) NULL
  );

ALTER TABLE
  public.customers
ADD
  CONSTRAINT customers_pkey PRIMARY KEY (customerid)