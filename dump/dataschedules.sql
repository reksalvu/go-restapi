--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-1.pgdg22.04+1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-1.pgdg22.04+1)

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
-- Data for Name: schedules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedules (id, title, start_at, end_at) FROM stdin;
1	Fitness	2022-06-01 00:00:00+07	2022-06-02 23:59:59+07
2	Learn	2022-06-05 00:00:00+07	2022-06-05 23:59:59+07
3	Working	2022-06-07 00:00:00+07	2022-06-12 23:59:59+07
\.


--
-- Name: schedules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedules_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--

