--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg110+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg110+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: atlas_schema_revisions; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA atlas_schema_revisions;


ALTER SCHEMA atlas_schema_revisions OWNER TO postgres;

--
-- Name: access_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.access_type AS ENUM (
    'read',
    'write'
);


ALTER TYPE public.access_type OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: atlas_schema_revisions; Type: TABLE; Schema: atlas_schema_revisions; Owner: postgres
--

CREATE TABLE atlas_schema_revisions.atlas_schema_revisions (
    version character varying NOT NULL,
    description character varying NOT NULL,
    type bigint DEFAULT 2 NOT NULL,
    applied bigint DEFAULT 0 NOT NULL,
    total bigint DEFAULT 0 NOT NULL,
    executed_at timestamp with time zone NOT NULL,
    execution_time bigint NOT NULL,
    error text,
    error_stmt text,
    hash character varying NOT NULL,
    partial_hashes jsonb,
    operator_version character varying NOT NULL
);


ALTER TABLE atlas_schema_revisions.atlas_schema_revisions OWNER TO postgres;

--
-- Name: access; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.access (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "user" uuid NOT NULL,
    type public.access_type NOT NULL,
    node uuid NOT NULL
);


ALTER TABLE public.access OWNER TO postgres;

--
-- Name: file; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.file (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    node uuid NOT NULL,
    content text NOT NULL
);


ALTER TABLE public.file OWNER TO postgres;

--
-- Name: node; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.node (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying NOT NULL,
    owner uuid NOT NULL,
    parent uuid
);


ALTER TABLE public.node OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Data for Name: atlas_schema_revisions; Type: TABLE DATA; Schema: atlas_schema_revisions; Owner: postgres
--

COPY atlas_schema_revisions.atlas_schema_revisions (version, description, type, applied, total, executed_at, execution_time, error, error_stmt, hash, partial_hashes, operator_version) FROM stdin;
20240411142653	user	2	1	1	2024-04-11 21:09:21.18252+00	176666			2IUYUyAiH2Ta0sgPI0OiG+JzQLXsthKkro/hqmZb+4Q=	["h1:86EE4dgIDdHdPJm05Cz9VvG+xY8kj27yqk33U+Ky3zM="]	Atlas CLI v0.21.2-e25c033-canary
20240411142713	node	2	1	1	2024-04-11 21:09:21.18622+00	198042			0QUkfQo+0sDgoNGvP+kSxqSN3KVmweg2S6W5r3adtvc=	["h1:vet/adCtkW/nfnqm8l+vxLq+JTtc81HNPOkT0vNN6IY="]	Atlas CLI v0.21.2-e25c033-canary
20240413000857		2	1	1	2024-04-13 00:54:37.247834+00	308209			2hfvR3ABvTLZfmC+0610AG7bzS7T+w0/l236wUPYvAc=	["h1:HW7pHIMJ+wXOsnQtqZX3xhnrqPpBcrpC1g2uKQHg7LU="]	Atlas CLI v0.21.2-e25c033-canary
20240413002501		2	2	2	2024-04-13 00:54:37.254697+00	244583			nwC/2Z8ibqx6ZoW1pRtaz2glXYKFFMxIq/6S2MNC59w=	["h1:hfd1OR3nfbc3NvU5qodRRUsTW2abMCh5Vv4ffnCxLKw=", "h1:e1AwpT/quUQPI8jf7nXO+/hxQbi8i52bh/o8B2/zYYc="]	Atlas CLI v0.21.2-e25c033-canary
20240413004158		2	2	2	2024-04-13 00:54:37.258228+00	164083			HR0LrbSMGjrmNB3sQQJ07BlhzhUzm0IStuHEzuQIdUU=	["h1:N7sxtFcgk6B9LQnKQgfw158pek6coSw2NpgUEyzX1i0=", "h1:57JLXVi+rdcvMbAUVHmMR7ydSfWMKM60OqSFrKnPfNQ="]	Atlas CLI v0.21.2-e25c033-canary
20240413013118		2	1	1	2024-04-13 05:27:54.639182+00	263125			kHdG94wI9T/0FflSC4a5aU5qCYPAXuabKvnciLcbOlU=	["h1:GflwZxG6dTxDxx4GcHd9nwWvsq1MEBknbnpzDJVvF/U="]	Atlas CLI v0.21.2-e25c033-canary
\.


--
-- Data for Name: access; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.access (id, "user", type, node) FROM stdin;
3071e2a5-abfd-474e-8ab3-24c4c415a3c0	3c403b9b-4303-4d51-b8f6-7218ef382745	read	01360ece-c6ed-4f69-83fc-8e3c5422d6e7
253a4fb7-5417-4c57-b377-0fb217817fc5	3c403b9b-4303-4d51-b8f6-7218ef382745	write	01360ece-c6ed-4f69-83fc-8e3c5422d6e7
\.


--
-- Data for Name: file; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.file (id, node, content) FROM stdin;
\.


--
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.node (id, name, owner, parent) FROM stdin;
01360ece-c6ed-4f69-83fc-8e3c5422d6e7	Foo	3c403b9b-4303-4d51-b8f6-7218ef382745	\N
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, name) FROM stdin;
3c403b9b-4303-4d51-b8f6-7218ef382745	Peter
\.


--
-- Name: atlas_schema_revisions atlas_schema_revisions_pkey; Type: CONSTRAINT; Schema: atlas_schema_revisions; Owner: postgres
--

ALTER TABLE ONLY atlas_schema_revisions.atlas_schema_revisions
    ADD CONSTRAINT atlas_schema_revisions_pkey PRIMARY KEY (version);


--
-- Name: access access_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access
    ADD CONSTRAINT access_pkey PRIMARY KEY (id);


--
-- Name: file file_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.file
    ADD CONSTRAINT file_pkey PRIMARY KEY (id);


--
-- Name: node node_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: access_user_type_node; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX access_user_type_node ON public.access USING btree ("user", type, node);


--
-- Name: node_name_owner_parent; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX node_name_owner_parent ON public.node USING btree (name, owner, parent);


--
-- Name: user_name_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_name_key ON public."user" USING btree (name);


--
-- Name: access access_node; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access
    ADD CONSTRAINT access_node FOREIGN KEY (node) REFERENCES public.node(id) ON DELETE CASCADE;


--
-- Name: access access_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access
    ADD CONSTRAINT access_user FOREIGN KEY ("user") REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: file file_node; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.file
    ADD CONSTRAINT file_node FOREIGN KEY (node) REFERENCES public.node(id) ON DELETE CASCADE;


--
-- Name: node node_owner; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_owner FOREIGN KEY (owner) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: node node_parent; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.node
    ADD CONSTRAINT node_parent FOREIGN KEY (parent) REFERENCES public.node(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

