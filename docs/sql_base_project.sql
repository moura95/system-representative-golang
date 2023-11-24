CREATE TABLE IF NOT EXISTS representative
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL,
    website      VARCHAR(255),
    logo_url     VARCHAR(255),
    street       VARCHAR(255),
    number       VARCHAR(30),
    city         VARCHAR(40),
    state        VARCHAR(255),
    zip_code     VARCHAR(255),
    is_active    BOOLEAN               DEFAULT TRUE NOT NULL,
    cnpj         VARCHAR(255) NOT NULL,
    fantasy_name VARCHAR(255) NOT NULL,
    ie           VARCHAR(255),
    phone        VARCHAR(255),
    phone2       VARCHAR(255),
    logo         VARCHAR(255),
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users
(
    id                SERIAL PRIMARY KEY NOT NULL,
    email             TEXT               NOT NULL UNIQUE,
    password          TEXT               NOT NULL,
    first_name      VARCHAR(255)        NOT NULL,
    last_name       VARCHAR(255)        NOT NULL,
    cpf             VARCHAR(255)        NOT NULL UNIQUE,
    is_active         BOOLEAN            NOT NULL DEFAULT TRUE,
    representative_id INT,
    created_at        TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP          NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE
);

-- Diretoria, Financeiro , Secretaria, Admin, Staff
CREATE TABLE IF NOT EXISTS permissions
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS users_permissions
(
    user_id       INT NOT NULL,
    permission_id INT NOT NULL,
    PRIMARY KEY (user_id, permission_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS representative_users
(
    representative_id INT NOT NULL,
    user_id           INT NOT NULL,
    PRIMARY KEY (representative_id, user_id),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "sessions"
(
    "id"            uuid PRIMARY KEY,
    "user_id"       int         NOT NULL,
    "refresh_token" varchar     NOT NULL,
    "user_agent"    varchar     NOT NULL,
    "client_ip"     varchar     NOT NULL,
    "is_blocked"    boolean     NOT NULL DEFAULT false,
    "expires_at"    timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS smtp
(
    user_id         INTEGER PRIMARY KEY NOT NULL,
    representative_id         INTEGER NOT NULL,
    email    VARCHAR(255),
    password VARCHAR(255),
    server VARCHAR(100),
    port VARCHAR(30),
    cryptography VARCHAR(20),
    created_at      TIMESTAMP           NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP           NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS sellers
(
    id                SERIAL PRIMARY KEY,
    representative_id INT          NOT NULL,
    name              VARCHAR(255) NOT NULL,
    pix               VARCHAR(255),
    email             VARCHAR(255) NOT NULL,
    phone             VARCHAR(255),
    phone2            VARCHAR(255),
    observation       VARCHAR(255),
    cpf               VARCHAR(255) NOT NULL,
    is_active         BOOLEAN               DEFAULT TRUE NOT NULL,
    created_at        TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP    NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE
);



CREATE TYPE company_types AS ENUM ('Factory', 'Customer', 'Portage');

CREATE TABLE IF NOT EXISTS companies
(
    id                SERIAL PRIMARY KEY,
    type              company_types NOT NULL,
    name              VARCHAR(255)  NOT NULL,
    email             VARCHAR(255)  NOT NULL,
    website           VARCHAR(255),
    logo_url          VARCHAR(255),
    street            VARCHAR(255),
    number            VARCHAR(30),
    city              VARCHAR(40),
    state             VARCHAR(255),
    zip_code          VARCHAR(255),
    is_active         BOOLEAN                DEFAULT TRUE NOT NULL,
    cnpj              VARCHAR(255)  NOT NULL,
    fantasy_name      VARCHAR(255)  NOT NULL,
    ie                VARCHAR(255),
    phone             VARCHAR(255),
    phone2            VARCHAR(255),
    created_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP     NOT NULL DEFAULT NOW(),
    representative_id INT           NOT NULL,
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE


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
    unidade           VARCHAR(255)   NOT NULL,
    description       VARCHAR(255),
    image_url         VARCHAR(255),
    is_active         BOOLEAN                 DEFAULT TRUE NOT NULL,
    created_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES companies (id) ON DELETE CASCADE
);



CREATE TYPE plan_types AS ENUM ('Basic', 'Advanced');

CREATE TABLE IF NOT EXISTS plans
(
    representative_id INT PRIMARY KEY,
    plan              plan_types NOT NULL,
    trial_period      boolean    NOT NULL DEFAULT TRUE,
    updated_at        TIMESTAMP  NOT NULL DEFAULT NOW(),
    data_expire       TIMESTAMP  NOT NULL DEFAULT NOW() + INTERVAL '7 days',
    created_at        TIMESTAMP  NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE
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



CREATE TYPE status_enum AS ENUM ('Rascunho','Cotacao','Cancelado','Concluido');

-- 15/30/45, 5/10/15 1/15, A vista, 30/60/90, 10/20/30, 7/14/21
CREATE TABLE IF NOT EXISTS form_payments
(
    id                SERIAL PRIMARY KEY,
    name              VARCHAR(255) NOT NULL,
    representative_id INT          NOT NULL,
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE
);

CREATE TYPE shipping_enum AS ENUM ('CIF','FOB','Outros');

CREATE TABLE IF NOT EXISTS orders
(
    id                SERIAL PRIMARY KEY,
    representative_id INT            NOT NULL,
    factory_id        INT            NOT NULL,
    customer_id       INT            NOT NULL,
    portage_id        INT            NOT NULL,
    seller_id         INT            NOT NULL,
    url_pdf           VARCHAR(255),
    buyer             VARCHAR(255),
    status            status_enum             DEFAULT 'Rascunho' NOT NULL,
    shipping          shipping_enum,
    form_payment_id   INT            NOT NULL,
    expire_order      TIMESTAMP      NOT NULL DEFAULT NOW() + INTERVAL '30 days',
    total             DECIMAL(10, 2) NOT NULL DEFAULT 0,
    is_active         BOOLEAN        NOT NULL DEFAULT true,
    created_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP      NOT NULL DEFAULT NOW(),
    FOREIGN KEY (representative_id) REFERENCES representative (id) ON DELETE CASCADE,
    FOREIGN KEY (factory_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (portage_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES sellers (id) ON DELETE CASCADE,
    FOREIGN KEY (form_payment_id) REFERENCES form_payments (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT            NOT NULL,
    product_id INT            NOT NULL,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    discount   DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
