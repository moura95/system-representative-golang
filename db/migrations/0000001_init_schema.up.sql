CREATE TYPE plan_types AS ENUM ('Trial', 'Silver', 'Gold');

CREATE TABLE IF NOT EXISTS representatives
(
    id           SERIAL PRIMARY KEY,
    cnpj         VARCHAR(255) UNIQUE,
    name         VARCHAR(255),
    fantasy_name VARCHAR(255),
    ie           VARCHAR(255),
    phone        VARCHAR(255),
    email        VARCHAR(255),
    website      VARCHAR(255),
    logo_url     VARCHAR(255),
    zip_code     VARCHAR(255),
    state        VARCHAR(255),
    city         VARCHAR(40),
    street       VARCHAR(255),
    number       VARCHAR(30),
    plan         plan_types NOT NULL DEFAULT 'Trial',
    stripe_id    VARCHAR(255),
    data_expire  TIMESTAMP  NOT NULL DEFAULT NOW() + INTERVAL '7 days',
    is_active    BOOLEAN    NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users
(
    id                SERIAL PRIMARY KEY NOT NULL,
    representative_id INT                NOT NULL,
    cpf               VARCHAR(255) UNIQUE,
    first_name        VARCHAR(255)       NOT NULL,
    last_name         VARCHAR(255)       NOT NULL,
    email             TEXT               NOT NULL UNIQUE,
    password          TEXT               NOT NULL,
    phone             VARCHAR(255),
    is_active         BOOLEAN            NOT NULL DEFAULT TRUE,
    last_login        TIMESTAMP          NOT NULL DEFAULT NOW(),
    created_at        TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP          NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS permissions
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_permissions
(
    user_id       INT NOT NULL,
    permission_id INT NOT NULL,
    PRIMARY KEY (user_id, permission_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sessions
(
    id                UUID PRIMARY KEY,
    user_id           INT         NOT NULL,
    representative_id INT         NOT NULL,
    refresh_token     VARCHAR     NOT NULL,
    user_agent        VARCHAR     NOT NULL,
    client_ip         VARCHAR     NOT NULL,
    is_blocked        BOOLEAN     NOT NULL DEFAULT false,
    expires_at        TIMESTAMPTZ NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sellers
(
    id                SERIAL PRIMARY KEY,
    representative_id INT          NOT NULL,
    cpf               VARCHAR(255) NOT NULL,
    name              VARCHAR(255) NOT NULL,
    phone             VARCHAR(255),
    email             VARCHAR(255),
    pix               VARCHAR(255),
    observation       VARCHAR(255),
    is_active         BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP    NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);

CREATE TYPE company_types AS ENUM ('Factory', 'Customer', 'Portage');

CREATE TABLE IF NOT EXISTS companies
(
    id                SERIAL PRIMARY KEY,
    representative_id INT           NOT NULL,
    type              company_types NOT NULL,
    cnpj              VARCHAR(255)  ,
    name              VARCHAR(255)  NOT NULL,
    fantasy_name      VARCHAR(255),
    ie                VARCHAR(255),
    phone             VARCHAR(255),
    email             VARCHAR(255),
    website           VARCHAR(255),
    logo_url          VARCHAR(255),
    zip_code          VARCHAR(255),
    state             VARCHAR(255),
    city              VARCHAR(40),
    street            VARCHAR(255),
    number            VARCHAR(30),
    is_active         BOOLEAN       NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);
CREATE TYPE payment_receipt_type AS ENUM ('Payment', 'Receipt');

create type payment_receipt_form_type as enum ('Pix', 'Invoice', 'Transfer', 'CreditCard', 'DebitCard', 'Cash', 'Cheque', 'Outros');

create type payment_receipt_status as enum ('Pending', 'Paid', 'Expired');

CREATE TABLE IF NOT EXISTS payment_receipt
(
    id                SERIAL PRIMARY KEY,
    representative_id INT           NOT NULL,
    type_payment              payment_receipt_type NOT NULL,
    status           payment_receipt_status NOT NULL DEFAULT 'Pending',
    description       VARCHAR(255)  NOT NULL,
    amount            DECIMAL(10, 2)  NOT NULL,
    expiration_date   TIMESTAMP,
    payment_date      TIMESTAMP,
    doc_number       VARCHAR(255),
    Recipient       VARCHAR(255),
    payment_form     payment_receipt_form_type NOT NULL DEFAULT 'Outros',
    is_active         BOOLEAN       NOT NULL DEFAULT TRUE,
    installment       INT NOT NULL DEFAULT 1,
    interval_days    INT NOT NULL DEFAULT 30,
    created_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    additional_info   VARCHAR(255),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS files_payment_receipt
(
    id                SERIAL PRIMARY KEY,
    payment_receipt_id INT           NOT NULL,
    url_file               VARCHAR(255)  NOT NULL,
    created_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (payment_receipt_id) REFERENCES payment_receipt (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS products
(
    id                SERIAL PRIMARY KEY,
    representative_id INT            NOT NULL,
    factory_id        INT            NOT NULL,
    name              VARCHAR(255)   NOT NULL,
    code              VARCHAR(255)   NOT NULL,
    price             DECIMAL(10, 2) NOT NULL,
    ipi               DECIMAL(10, 2),
    reference         VARCHAR(255),
    description       VARCHAR(255),
    image_url         VARCHAR(255),
    is_active         BOOLEAN        NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES companies (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS form_payments
(
    id                SERIAL PRIMARY KEY,
    representative_id INT          NOT NULL,
    name              VARCHAR(255) NOT NULL,
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);

CREATE TYPE status_enum AS ENUM ('Rascunho','Cotacao','Cancelado','Concluido');

CREATE TYPE shipping_enum AS ENUM ('CIF','FOB','Outros');

CREATE TABLE IF NOT EXISTS orders
(
    id                SERIAL PRIMARY KEY,
    representative_id INT            NOT NULL,
    factory_id        INT            NOT NULL,
    customer_id       INT            NOT NULL,
    portage_id        INT            NOT NULL,
    seller_id         INT            NOT NULL,
    form_payment_id   INT            ,
    order_number      INT            NOT NULL,
    url_pdf           VARCHAR(255),
    buyer             VARCHAR(255),
    shipping          shipping_enum  NOT NULL DEFAULT 'Outros',
    status            status_enum    NOT NULL DEFAULT 'Rascunho',
    expired_at        TIMESTAMP      NOT NULL DEFAULT NOW() + INTERVAL '30 days',
    total             DECIMAL(10, 2) NOT NULL DEFAULT 0,
    is_active         BOOLEAN        NOT NULL DEFAULT true,
    created_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (portage_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES sellers (id) ON DELETE CASCADE,
    FOREIGN KEY (form_payment_id) REFERENCES form_payments (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS files_orders
(
    id                SERIAL PRIMARY KEY,
    order_id INT           NOT NULL,
    url_file               VARCHAR(255)  NOT NULL,
    created_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE);

CREATE TABLE IF NOT EXISTS order_items
(
    order_id   INT            NOT NULL,
    product_id INT            NOT NULL,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    discount   DECIMAL(10, 2) not null default 0,
    total      DECIMAL(10, 2) NOT NULL DEFAULT 0,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS calendars
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255)          NOT NULL,
    visit_start TIMESTAMP             NOT NULL,
    visit_end   TIMESTAMP             NOT NULL,
    allday      BOOLEAN DEFAULT FALSE NOT NULL,
    user_id     INT                   NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS smtp
(
    representative_id INTEGER PRIMARY KEY NOT NULL,
    is_active         BOOLEAN                      DEFAULT TRUE NOT NULL,
    email             VARCHAR(255)        NOT NULL,
    password          VARCHAR(255)        NOT NULL,
    server            VARCHAR(100)        NOT NULL,
    port              INTEGER             NOT NULL,
    created_at        TIMESTAMP           NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP           NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);


CREATE TYPE origin_leads_enum AS ENUM ('Facebook','Instagram','Google','Linkedin','Outros');

CREATE TABLE IF NOT EXISTS leads
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)      NOT NULL,
    email      VARCHAR(255)      NOT NULL unique,
    phone      VARCHAR(255)      NOT NULL,
    origin     origin_leads_enum NOT NULL DEFAULT 'Outros',
    is_active  BOOLEAN           NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP         NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP         NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS activity
(
    id          SERIAL PRIMARY KEY,
    action        VARCHAR(20) NOT NULL,
    reference_url VARCHAR(255) NOT NULL,
    user_id       INT          NOT NULL,
    representative_id INT     NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (representative_id) REFERENCES representatives (id) ON DELETE CASCADE
);

INSERT INTO permissions (name)
VALUES ('Admin');

INSERT INTO permissions (name)
VALUES ('Staff');

INSERT INTO permissions (name)
VALUES ('Admin_Representative');

INSERT INTO permissions (name)
VALUES ('Director');

INSERT INTO permissions (name)
VALUES ('Financial');

INSERT INTO permissions (name)
VALUES ('Secretary');

INSERT INTO permissions (name)
VALUES ('Seller');

CREATE OR REPLACE FUNCTION update_order_total()
    RETURNS TRIGGER AS
$$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        UPDATE orders
        SET total = (SELECT COALESCE(SUM(((price * quantity) * (1 - discount / 100))), 0)
                     FROM order_items
                     WHERE order_id = old.order_id)
        WHERE id = old.order_id;
    ELSE
        UPDATE orders
        SET total = (SELECT COALESCE(SUM(((price * quantity) * (1 - discount / 100))), 0)
                     FROM order_items
                     WHERE order_id = new.order_id)
        WHERE id = new.order_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_order_total_trigger
    AFTER INSERT OR UPDATE OR DELETE
    ON order_items
    FOR EACH ROW
EXECUTE FUNCTION update_order_total();

CREATE OR REPLACE FUNCTION update_order_items_total()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.total = (NEW.quantity * NEW.price) * (1 - NEW.discount / 100);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER order_items_total_trigger
    BEFORE INSERT OR UPDATE
    ON order_items
    FOR EACH ROW
EXECUTE FUNCTION update_order_items_total();

SET TIME ZONE 'America/Sao_Paulo';


insert into representatives (cnpj, name, fantasy_name, email, phone, ie, website, logo_url, zip_code, state, city,
                             street, number, plan)
values ('90669024000120', 'Midas', 'Midas LTDA', 'midas@hotmail.com', '9199595959', '514545', '', '',
        '66000-000', 'PA', 'Belem', 'Avenida Santo amaro', '453', 'Gold');

insert into users (representative_id, cpf, first_name, last_name, email, password, phone)
values (1, '0261133', 'Tarcisio', 'Lucas', 'dare-test@hotmail.com',
        '$2a$10$K.OETkLj7ARmNB2fsqA7PO1zXorL6qYrCGSbIhDG0qilL4sMNJShq', '9199595959');

insert into users (representative_id, cpf, first_name, last_name, email, password, phone)
values (1, '02651132293', 'Junior', 'Moura', 'jr@hotmail.com',
        '$2a$10$K.OETkLj7ARmNB2fsqA7PO1zXorL6qYrCGSbIhDG0qilL4sMNJShq', '9199895959');

insert into users (representative_id, cpf, first_name, last_name, email, password, phone)
values (1, '02651132294', 'Ivan', 'Noleto', 'ivannoleto@hotmail.com',
        '$2a$10$K.OETkLj7ARmNB2fsqA7PO1zXorL6qYrCGSbIhDG0qilL4sMNJShq', '9199895957');

insert into form_payments(name, representative_id)
values ('A Vista', 1);

insert into form_payments(name, representative_id)
values ('5/10/15', 1);

insert into form_payments(name, representative_id)
values ('10/20/30', 1);

insert into form_payments(name, representative_id)
values ('15/30/45', 1);

insert into form_payments(name, representative_id)
values ('30/60/90', 1);

INSERT INTO user_permissions (user_id, permission_id)
VALUES (1, 1);

INSERT INTO user_permissions (user_id, permission_id)
VALUES (2, 1);

INSERT INTO user_permissions (user_id, permission_id)
VALUES (3, 1);

insert into sellers (representative_id, cpf, name, email, phone)
values (1, '90669024000120', 'Luiz Moureli', 'junior.moura19@hotmail.com', '91983225598');

insert into sellers (representative_id, cpf, name, email, phone)
values (1, '9066902550120', 'Pedro Moureli', 'junior.moura19@hotmail.com', '91983225598');

insert into companies (representative_id, type, name, cnpj)
values (1, 'Portage', 'Outros', '000000000');

insert into companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj, fantasy_name, ie, phone)
values (1, 'Portage', 'Transportadora BR', 'vendas@br.com.br', 'www.transbr.com.br', '', 'Rua Augusto Montenegro',
        '543', 'Sao Paulo', 'SP', '04317-000', '90669024000120', 'Trans BR LTDA', '367920860324',
        '91 984915598');

insert into companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj, fantasy_name, ie, phone)
values (1, 'Factory', 'Cleber', 'vendas@cleber.com.br', 'www.cleber.com.br', '', 'Rua Santarem', '123', 'Sao Paulo',
        'SP', '04317-000', '90669024000120', 'Cleber LTDA', '367920860324', '91 983225598');

insert into companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj, fantasy_name, ie, phone)
values (1, 'Factory', 'BMW', 'vendas@bmw.com.br', 'www.bmw.com.br', '', 'Rua Santarem', '123', 'Sao Paulo',
        'SP', '04317-000', '9066902100120', 'BMW LTDA', '367920860124', '91 983225598');

insert into companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj, fantasy_name, ie, phone)
values (1, 'Customer', 'EletroHidro', 'vendas@eletrohidro.com.br', 'www.eletrohidro.com.br', '', 'Rua Magalhaes barata',
        '123', 'Belem', 'PA', '66615-000', '90669028700120', 'Eletro hidro LTDA', '367920860324',
        '91 984915598');

insert into companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj, fantasy_name, ie, phone)
values (1, 'Customer', 'ENEL', 'vendas@enel.com.br', 'www.enel.com.br', '', 'Rua Magalhaes barata',
        '123', 'Belem', 'PA', '66615-000', '90669028700121', 'Eletro hidro LTDA', '36920860324',
        '91 984915598');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 3, 'Prego Telheiro', 'pg5512', '5.50', '5.0', 'Prego Telheiro 12x50', '');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 3, 'Prego Feio', 'pg5513', '7.50', '5.0', 'Prego Telheiro 12x70', '');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 3, 'Casa', '5514', '8.50', '5.0', 'Prego Telheiro 12x80', '');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 4, 'M1', '5514', '200000', '5.0', 'Carro', '');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 4, 'M3', '5514', '500000', '5.0', 'Carro', '');

insert into products(representative_id, factory_id, name, code, price, ipi, reference, image_url)
values (1, 4, 'M7', '5514', '10000000', '5.0', 'Carro', '');

insert into orders(representative_id, factory_id, customer_id, portage_id, seller_id, order_number,
                   url_pdf, buyer, status, shipping, form_payment_id, created_at)
values (1, 4, 5, 4, 1, 1, '', 'Jose da Silva', 'Concluido', 'CIF', 1, '2023-02-11');

insert into orders(representative_id, factory_id, customer_id, portage_id, seller_id, order_number,
                   url_pdf, buyer, status, shipping, form_payment_id, created_at)
values (1, 4, 5, 4, 1, 2, '', 'Jose da Silva', 'Concluido', 'CIF', Null, '2023-02-10');

insert into orders(representative_id, factory_id, customer_id, portage_id, seller_id, order_number,
                   url_pdf, buyer, status, shipping, form_payment_id)
values (1, 4, 6, 4, 1, 3, '', 'Jose da Silva', 'Rascunho', 'CIF', 2);

insert into order_items(order_id, product_id, quantity, price, discount)
values (1, 1, 10, 5.50, 10.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (1, 2, 20, 7.50, 25.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (1, 3, 10, 8.50, 0.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (2, 4, 10, 5.50, 10.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (2, 5, 10, 8.50, 0.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (3, 4, 10, 5.50, 10.0);

insert into order_items(order_id, product_id, quantity, price, discount)
values (3, 1, 10, 8.50, 0.0);
