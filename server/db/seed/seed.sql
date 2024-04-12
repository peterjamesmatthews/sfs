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
\.


--
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.node (id, name, owner, parent) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, name) FROM stdin;
\.


--
-- Name: atlas_schema_revisions atlas_schema_revisions_pkey; Type: CONSTRAINT; Schema: atlas_schema_revisions; Owner: postgres
--

ALTER TABLE ONLY atlas_schema_revisions.atlas_schema_revisions
    ADD CONSTRAINT atlas_schema_revisions_pkey PRIMARY KEY (version);


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

