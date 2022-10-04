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
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.activities (id, schedule_id, activity) FROM stdin;
1	1	Elit reprehenderit ex ut
2	1	Exercitation ex sit cillum
3	2	Culpa ad esse sint veniam
4	2	Aliqua qui dolor do anim ex.
5	3	Eiusmod laborum proident
6	3	Officia occaecat
7	3	Sit aute commodo
8	3	Sit aute commodoEt
9	3	Excepteur proident sunt
\.


--
-- Name: activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.activities_id_seq', 9, true);


--
-- PostgreSQL database dump complete
--

