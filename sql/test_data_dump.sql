--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.1
-- Dumped by pg_dump version 9.5.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: post; Type: TABLE; Schema: public; Owner: lee
--

CREATE TABLE post (
    id integer NOT NULL,
    date_posted timestamp without time zone DEFAULT now(),
    subject text,
    text text,
    thread_id integer NOT NULL
);


ALTER TABLE post OWNER TO lee;

--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: lee
--

CREATE SEQUENCE post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE post_id_seq OWNER TO lee;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lee
--

ALTER SEQUENCE post_id_seq OWNED BY post.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: lee
--

ALTER TABLE ONLY post ALTER COLUMN id SET DEFAULT nextval('post_id_seq'::regclass);


--
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: lee
--

COPY post (id, date_posted, subject, text, thread_id) FROM stdin;
16	2016-11-05 12:20:13.065077	Thread 1	Thread one post - 1	1
17	2016-11-05 12:20:16.256236	Thread 1	Thread one post - 2	1
18	2016-11-05 12:20:18.176262	Thread 1	Thread one post - 3	1
22	2016-11-05 12:20:31.704861	Thread 1	Thread one post - 4	1
19	2016-11-05 12:20:24.064493	Thread 2	Thread two post - 1	2
20	2016-11-05 12:20:25.576404	Thread 2	Thread two post - 2	2
21	2016-11-05 12:20:27.112532	Thread 2	Thread two post - 3	2
\.


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: lee
--

SELECT pg_catalog.setval('post_id_seq', 22, true);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

