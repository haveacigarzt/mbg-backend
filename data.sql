--
-- PostgreSQL database dump
--

\restrict YIGuEs4fMvYa0rwCQjyqwIH3LPqeu8SGIzxVS9aSRmfHBDnU0Qm5SRqgSP5TiIA

-- Dumped from database version 18.3
-- Dumped by pg_dump version 18.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: kecamatan; Type: TABLE DATA; Schema: public; Owner: rian
--

SET SESSION AUTHORIZATION DEFAULT;

ALTER TABLE public.kecamatan DISABLE TRIGGER ALL;

INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (1, '2026-05-24 22:58:08+07', 'TOBA', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (2, '2026-05-24 22:58:08+07', 'MELIAU', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (3, '2026-05-24 22:58:08+07', 'KAPUAS', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (4, '2026-05-24 22:58:08+07', 'MUKOK', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (5, '2026-05-24 22:58:08+07', 'JANGKANG', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (6, '2026-05-24 22:58:08+07', 'BONTI', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (7, '2026-05-24 22:58:08+07', 'PARINDU', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (8, '2026-05-24 22:58:08+07', 'TAYAN HILIR', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (9, '2026-05-24 22:58:08+07', 'BALAI', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (10, '2026-05-24 22:58:08+07', 'TAYAN HULU', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (11, '2026-05-24 22:58:08+07', 'KEMBAYAN', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (13, '2026-05-24 22:58:08+07', 'NOYAN', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (14, '2026-05-24 22:58:08+07', 'SEKAYAM', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (15, '2026-05-24 22:58:08+07', 'ENTIKONG', 1);
INSERT INTO public.kecamatan (id, created_at, name, version) VALUES (12, '2026-05-24 22:58:08+07', 'BEDUAI', 2);


ALTER TABLE public.kecamatan ENABLE TRIGGER ALL;

--
-- Data for Name: kelurahan; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.kelurahan DISABLE TRIGGER ALL;

INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (1, '2026-05-24 23:07:22+07', 1, 'BAGAN ASAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (2, '2026-05-24 23:07:22+07', 1, 'TERAJU BARAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (3, '2026-05-24 23:07:22+07', 1, 'KAMPUNG BARU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (4, '2026-05-24 23:07:22+07', 1, 'SANSAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (5, '2026-05-24 23:07:22+07', 1, 'BALAI BELUNGAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (6, '2026-05-24 23:07:22+07', 1, 'BELUNGAI DALAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (7, '2026-05-24 23:07:22+07', 2, 'BARU LOMBAK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (8, '2026-05-24 23:07:22+07', 2, 'KUNYIL', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (9, '2026-05-24 23:07:22+07', 2, 'PAMPANG DUA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (10, '2026-05-24 23:07:22+07', 2, 'HARAPAN MAKMUR', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (11, '2026-05-24 23:07:22+07', 2, 'SUNGAI KEMBAYAU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (12, '2026-05-24 23:07:22+07', 2, 'KUALA ROSAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (13, '2026-05-24 23:07:22+07', 2, 'KUALA BUAYAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (14, '2026-05-24 23:07:22+07', 2, 'BAKTI JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (15, '2026-05-24 23:07:22+07', 2, 'CUPANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (16, '2026-05-24 23:07:22+07', 2, 'MUKTI JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (17, '2026-05-24 23:07:22+07', 2, 'LALANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (18, '2026-05-24 23:07:22+07', 2, 'ENGGADAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (19, '2026-05-24 23:07:22+07', 2, 'MERANGGAU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (20, '2026-05-24 23:07:22+07', 2, 'BALAI TINGGI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (21, '2026-05-24 23:07:22+07', 2, 'MELIAU HILIR', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (22, '2026-05-24 23:07:22+07', 2, 'MELIAU HULU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (23, '2026-05-24 23:07:22+07', 2, 'SUNGAI MAYAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (24, '2026-05-24 23:07:22+07', 2, 'MELOBOK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (25, '2026-05-24 23:07:22+07', 2, 'MELAWI MAKMUR', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (26, '2026-05-24 23:07:22+07', 3, 'PENYALIMAU JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (27, '2026-05-24 23:07:22+07', 3, 'PENYALIMAU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (28, '2026-05-24 23:07:22+07', 3, 'RAMBIN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (29, '2026-05-24 23:07:22+07', 3, 'NANGA BIANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (30, '2026-05-24 23:07:22+07', 3, 'LINTANG PELAMAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (31, '2026-05-24 23:07:22+07', 3, 'SEI ALAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (32, '2026-05-24 23:07:22+07', 3, 'SEMERANGKAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (33, '2026-05-24 23:07:22+07', 3, 'SUNGAI BATU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (34, '2026-05-24 23:07:22+07', 3, 'SUNGAI MUNTIK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (35, '2026-05-24 23:07:22+07', 3, 'LINTANG KAPUAS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (36, '2026-05-24 23:07:22+07', 3, 'BELANGIN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (37, '2026-05-24 23:07:22+07', 3, 'PENYELADI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (38, '2026-05-24 23:07:22+07', 3, 'TANJUNG KAPUAS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (39, '2026-05-24 23:07:22+07', 3, 'TANJUNG SEKAYAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (40, '2026-05-24 23:07:22+07', 3, 'ILIR KOTA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (41, '2026-05-24 23:07:22+07', 3, 'BERINGIN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (42, '2026-05-24 23:07:22+07', 3, 'LAPE', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (43, '2026-05-24 23:07:22+07', 3, 'SEI MAWANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (44, '2026-05-24 23:07:22+07', 3, 'SEI SENGKUANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (45, '2026-05-24 23:07:22+07', 3, 'PANA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (46, '2026-05-24 23:07:22+07', 3, 'MANGKIANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (47, '2026-05-24 23:07:22+07', 3, 'ENTAKAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (48, '2026-05-24 23:07:22+07', 3, 'KAMBONG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (49, '2026-05-24 23:07:22+07', 3, 'TAPANG DULANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (50, '2026-05-24 23:07:22+07', 3, 'BOTUH LINTANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (51, '2026-05-24 23:07:22+07', 4, 'INGGIS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (52, '2026-05-24 23:07:22+07', 4, 'SEMANGGIS RAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (53, '2026-05-24 23:07:22+07', 4, 'SEMUNTAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (54, '2026-05-24 23:07:22+07', 4, 'ENGKODE', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (55, '2026-05-24 23:07:22+07', 4, 'SEI MAWANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (56, '2026-05-24 23:07:22+07', 4, 'TRIMULYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (57, '2026-05-24 23:07:22+07', 4, 'LAYAK OMANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (58, '2026-05-24 23:07:22+07', 4, 'SERAMBAI JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (59, '2026-05-24 23:13:50+07', 5, 'TERATI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (60, '2026-05-24 23:13:50+07', 5, 'SELAMPUNG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (61, '2026-05-24 23:13:50+07', 5, 'SAPE', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (62, '2026-05-24 23:13:50+07', 5, 'SEMIRAU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (63, '2026-05-24 23:13:50+07', 5, 'BALAI SEBUT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (64, '2026-05-24 23:13:50+07', 5, 'SEMOMBAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (65, '2026-05-24 23:13:50+07', 5, 'EMPIYANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (66, '2026-05-24 23:13:50+07', 5, 'JANGKANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (67, '2026-05-24 23:13:50+07', 5, 'TANGGUNG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (68, '2026-05-24 23:13:50+07', 5, 'PISANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (69, '2026-05-24 23:13:50+07', 5, 'KETORI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (70, '2026-05-24 23:13:50+07', 6, 'BAHTA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (71, '2026-05-24 23:13:50+07', 6, 'TUNGGUL BOYOK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (72, '2026-05-24 23:13:50+07', 6, 'SAMI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (73, '2026-05-24 23:13:50+07', 6, 'EMPODIS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (74, '2026-05-24 23:13:50+07', 6, 'BONTI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (75, '2026-05-24 23:13:50+07', 6, 'KAMPUH', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (76, '2026-05-24 23:13:50+07', 6, 'BANTAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (77, '2026-05-24 23:13:50+07', 6, 'MAJEL', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (78, '2026-05-24 23:13:50+07', 7, 'MARITA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (79, '2026-05-24 23:13:50+07', 7, 'EMBALA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (80, '2026-05-24 23:13:50+07', 7, 'PANDU RAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (81, '2026-05-24 23:13:50+07', 7, 'MAJU KARYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (82, '2026-05-24 23:13:50+07', 7, 'GUNAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (83, '2026-05-24 23:13:50+07', 7, 'SUKA GRUNDI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (84, '2026-05-24 23:13:50+07', 7, 'SUKA MULYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (85, '2026-05-24 23:13:50+07', 7, 'PALEM JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (86, '2026-05-24 23:13:50+07', 7, 'PUSAT DAMAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (87, '2026-05-24 23:13:50+07', 7, 'SEBARA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (88, '2026-05-24 23:13:50+07', 7, 'HIBUN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (89, '2026-05-24 23:13:50+07', 7, 'RAHAYU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (90, '2026-05-24 23:13:50+07', 7, 'MARIGIN JAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (91, '2026-05-24 23:13:50+07', 7, 'DOSAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (92, '2026-05-24 23:13:50+07', 8, 'LALANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (93, '2026-05-24 23:13:50+07', 8, 'KAWAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (94, '2026-05-24 23:13:50+07', 8, 'PULAU TAYAN UTARA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (95, '2026-05-24 23:13:50+07', 8, 'PEDALAMAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (96, '2026-05-24 23:13:50+07', 8, 'TANJUNG BUNUT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (97, '2026-05-24 23:13:50+07', 8, 'SEBEMBAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (98, '2026-05-24 23:13:50+07', 8, 'BEGINJAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (99, '2026-05-24 23:13:50+07', 8, 'SUNGAI JAMAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (100, '2026-05-24 23:13:50+07', 8, 'EMBERAS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (101, '2026-05-24 23:13:50+07', 8, 'MELUGAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (102, '2026-05-24 23:13:50+07', 8, 'CEMPEDAK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (103, '2026-05-24 23:13:50+07', 8, 'SEJOTANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (104, '2026-05-24 23:13:50+07', 8, 'SUBAH', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (105, '2026-05-24 23:13:50+07', 8, 'TEBANG BENUA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (106, '2026-05-24 23:13:50+07', 8, 'BALAI INGIN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (107, '2026-05-24 23:13:50+07', 9, 'SEMONCOL', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (108, '2026-05-24 23:13:50+07', 9, 'MAK KAWING', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (109, '2026-05-24 23:13:50+07', 9, 'BULU BALA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (110, '2026-05-24 23:13:50+07', 9, 'TEMIANG TABA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (111, '2026-05-24 23:13:50+07', 9, 'SENYABANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (112, '2026-05-24 23:13:50+07', 9, 'KEBADU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (113, '2026-05-24 23:13:50+07', 9, 'HILIR', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (114, '2026-05-24 23:13:50+07', 9, 'TEMIANG MALI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (115, '2026-05-24 23:13:50+07', 9, 'TAE', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (116, '2026-05-24 23:13:50+07', 9, 'PADI KAYE', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (117, '2026-05-24 23:13:50+07', 9, 'EMPIRANG UJUNG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (118, '2026-05-24 23:13:50+07', 10, 'MENYABO', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (119, '2026-05-24 23:13:50+07', 10, 'BINJAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (120, '2026-05-24 23:13:50+07', 10, 'PANDAN SEMBUAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (121, '2026-05-24 23:13:50+07', 10, 'KEDAKAS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (122, '2026-05-24 23:13:50+07', 10, 'SOSOK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (123, '2026-05-24 23:13:50+07', 10, 'PERUAN DALAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (124, '2026-05-24 23:13:50+07', 10, 'MANDONG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (125, '2026-05-24 23:13:50+07', 10, 'JANJANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (126, '2026-05-24 23:13:50+07', 10, 'RIYAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (127, '2026-05-24 23:13:50+07', 10, 'BERAKAK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (128, '2026-05-24 23:13:50+07', 10, 'ENGKASAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (129, '2026-05-24 23:21:53+07', 11, 'KELOMPU', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (130, '2026-05-24 23:21:53+07', 11, 'TANAP', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (131, '2026-05-24 23:21:53+07', 11, 'MOBUI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (132, '2026-05-24 23:21:53+07', 11, 'SEJUAH', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (133, '2026-05-24 23:21:53+07', 11, 'TANJUNG MERPATI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (134, '2026-05-24 23:21:53+07', 11, 'SEBONGKUH', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (135, '2026-05-24 23:21:53+07', 11, 'KUALA DUA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (136, '2026-05-24 23:21:53+07', 11, 'TUNGGAL BHAKTI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (137, '2026-05-24 23:21:53+07', 11, 'SEMAYANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (138, '2026-05-24 23:21:53+07', 11, 'TANJUNG BUNGA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (139, '2026-05-24 23:21:53+07', 12, 'SEI ILAI', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (140, '2026-05-24 23:21:53+07', 12, 'BERENG BERKAWAT', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (141, '2026-05-24 23:21:53+07', 12, 'KASROMEGO', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (142, '2026-05-24 23:21:53+07', 12, 'THANG RAYA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (143, '2026-05-24 23:21:53+07', 12, 'MAWANG MUDA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (144, '2026-05-24 23:21:53+07', 13, 'EMPOTO', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (145, '2026-05-24 23:21:53+07', 13, 'IDAS', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (146, '2026-05-24 23:21:53+07', 13, 'SEI DANGIN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (147, '2026-05-24 23:21:53+07', 13, 'SEMONGAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (148, '2026-05-24 23:21:53+07', 13, 'NOYAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (149, '2026-05-24 23:21:53+07', 14, 'SOTOK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (150, '2026-05-24 23:21:53+07', 14, 'PENGADANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (151, '2026-05-24 23:21:53+07', 14, 'KENAMAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (152, '2026-05-24 23:21:53+07', 14, 'RAUT MUARA', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (153, '2026-05-24 23:21:53+07', 14, 'ENGKAHAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (154, '2026-05-24 23:21:53+07', 14, 'BALAI KARANGAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (155, '2026-05-24 23:21:53+07', 14, 'BUNGKANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (156, '2026-05-24 23:21:53+07', 14, 'LUBUK SABUK', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (157, '2026-05-24 23:21:53+07', 14, 'MALENGGANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (158, '2026-05-24 23:21:53+07', 14, 'SEI TEKAM', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (159, '2026-05-24 23:21:53+07', 15, 'NEKAN', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (160, '2026-05-24 23:21:53+07', 15, 'SEMANGET', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (161, '2026-05-24 23:21:53+07', 15, 'ENTIKONG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (162, '2026-05-24 23:21:53+07', 15, 'PALAPASANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (163, '2026-05-24 23:21:53+07', 15, 'SURUH TEMBAWANG', 1);
INSERT INTO public.kelurahan (id, created_at, kecamatan_id, name, version) VALUES (164, '2026-06-01 16:38:59+07', 3, 'BUNUT', 1);


ALTER TABLE public.kelurahan ENABLE TRIGGER ALL;

--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.roles DISABLE TRIGGER ALL;

INSERT INTO public.roles (id, name, permissions) VALUES (1, 'admin', '{akunsppg:write,akunsppg:read,sppg:read,sekolah:read,posyandu:read,pengiriman:read,tracking:read}');
INSERT INTO public.roles (id, name, permissions) VALUES (2, 'stakeholder', '{sppg:read,sekolah:read,posyandu:read,pengiriman:read,tracking:read}');
INSERT INTO public.roles (id, name, permissions) VALUES (3, 'sppg', '{sppg:read,sppg:write,sekolah:read,sekolah:write,posyandu:read,posyandu:write,pengiriman:read,tracking:read,driver:read,driver:write,pengiriman:write}');
INSERT INTO public.roles (id, name, permissions) VALUES (4, 'driver', '{pengiriman:read,tracking:read,tracking:write,pengiriman:write}');


ALTER TABLE public.roles ENABLE TRIGGER ALL;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.users DISABLE TRIGGER ALL;

INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (75, '2026-07-02 20:34:02+07', 'Dinas XXX', 'dinasxxx@gmail.com', '\x24326124313224784b6f6731326b664e6f4c74546c42394352516933654935596755547541512e4d512f49634f4e4842444845784b33346d476f6753', true, 1, 2, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (72, '2026-06-26 20:21:38+07', 'sppg_pontianak', 'admin@sppg.example.com', '\x243261243132246b7553496636795553785a707756346535724a4b6f6558717876634e6f655a646232374238426b63564356384a5369507771415061', true, 1, 3, '2026-06-26 23:53:45.60979', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (79, '2026-07-02 20:53:47+07', 'SPPG Test2', 'sppg_test2@gmail.com', '\x2432612431322477535175737636633076543450745a435a794a5152754d7272547067527171394676387375574f4730784f30314e3938424438302e', true, 1, 3, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (56, '2026-06-24 18:21:58+07', 'PRAMANA AFRIANDY', 'pramanaafriandy@gmail.com', '\x243261243132247833583569395338656b414a6e6e7a312f6b6d4d6375522e55797a632f39524c626f70584178536842773739577a4a416a6e775132', true, 1, 3, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (34, '2026-05-26 11:15:10+07', 'Admin SPPG Pontianak Barat', 'admin_ponbar@sppg.id', '\x2432612431322472435442716e4f5a4d524263562f444c366e616b6c7544587a3842735949306b4f706d612f41505144624e62496c7876676c753332', true, 1, 3, '2026-07-05 22:49:10.233053', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (47, '2026-05-27 18:34:21+07', 'Driver Bunut 2', 'driver.bunut2@gmail.com', '\x2432612431322459316c58493564495a785a4b6438755261653759686537386c43512e6f3159312f552e75734a51303032534a5a72566b666c507753', true, 1, 4, '2026-07-06 02:16:01.707056', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (73, '2026-07-02 20:32:34+07', 'SPPG Test', 'sppg_test@gmail.com', '\x2432612431322478633959774a6f305a344f396f42492e775077346e756e616149485a3337686e3644556d2e59596570636253712e684a6c2e6c4269', false, 3, 3, NULL, '2026-07-05 02:01:10.602513');
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (48, '2026-06-01 09:22:42+07', 'Dinas Sanggau', 'dinassgu@gmail.com', '\x24326124313224774a675874413644444c53596971736d346e646c5365696e3049624958392e667839414e6c6e304247456a75364e536b7a6b307775', false, 1, 2, '2026-07-07 21:07:22.731266', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (50, '2026-06-01 11:02:48+07', 'Admin SPPG Pontianak Timur', 'test@gmail.com', '\x243261243132242e774657386b755a4a352f6f75514539633955525a65384a587a5a674f4c49704a71416f773677625a6779306173764d2e73516169', true, 1, 3, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (52, '2026-06-01 11:03:56+07', 'Admin SPPG Pontianak Timur', 'test@gmail1.com', '\x24326124313224484b66705a374659596a362f376f505a65395237574f7950774f33374a34626a4e6e714871696251706744316e644741717a4c782e', true, 1, 3, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (54, '2026-06-01 11:05:37+07', 'Admin SPPG Pontianak Timur', 'test@gmail2.com', '\x243261243132244a4f642f6a616979493243414d414668316c4469674f456a50425569586e4f50447133572f5073422e696d5847526247413845714b', true, 1, 3, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (87, '2026-07-02 22:54:38+07', 'Test 2', 'admin@sppg.example.co', '\x2432612431322465784a64773972345955346e624c4c754e4f4172314f5a5a536d6a7353516b3679746c783172746863566e55483058333871536c6d', true, 3, 3, '2026-07-04 01:20:40.942624', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (55, '2026-06-02 16:21:13+07', 'Driver Bunut 1', 'driver.bunut1@gmail.com', '\x243261243132244a3043746e6f4b42483753354733617930444e702f755141524b5935504c494a696a4a775a544832456155356e63513639666b652e', true, 1, 4, NULL, NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (39, '2026-05-26 19:06:45+07', 'Admin SPPG Pontianak Timur', 'admin_pontim@sppg.id', '\x24326124313224322e4d637370634c2f3348584b464f4e61562f55586574505843536d6e544c68585479586c6a656353303738554a6f777066702f69', true, 1, 3, '2026-07-13 11:55:55.369026', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (37, '2026-05-26 11:34:30+07', 'Rian', 'afriandy193@gmail.com', '\x243261243132246b633732364948335970746945486f4e696556374f6562785230552f70755134656a4e442f446d39474d773336742e52706675414b', true, 2, 1, '2026-07-12 21:11:36.014695', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (46, '2026-05-27 17:37:41+07', 'Driver Ilir Kota', 'driver.ilirkota@gmail.com', '\x243261243132246b7043416c6c327670555178444a564f3851586c4e6535367448536232512f5267546657616546535a706d366469374b6661745061', true, 1, 4, '2026-07-12 21:13:13.902212', NULL);
INSERT INTO public.users (id, created_at, name, email, password_hash, activated, version, role_id, last_active_at, deleted_at) VALUES (71, '2026-06-24 20:46:23+07', 'PRAMANA AFRIANDY', 'pramanaafriandy@gmail.id', '\x243261243132246c5a394362444c51496d557750532e775a5933784e754363664d7573586d6f424753634a6c6e71584f6a63737034755451736c6871', true, 2, 3, NULL, '2026-07-04 19:58:43.716467');


ALTER TABLE public.users ENABLE TRIGGER ALL;

--
-- Data for Name: sppg; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.sppg DISABLE TRIGGER ALL;

INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (5, 39, '2026-05-26 19:06:45+07', '2026-05-26 19:06:45+07', 'SPPG ILIR KOTA - KOTA SANGGAU', 'Jl. Jenderal Ahmad Yani No.17', '{https://instagram.com/sppgsanggaukpsilirkota}', 'Siti Rahmawati', '081345678901', 'sppg.kapuas.ilirkota@email.com', 0.1235888731544522, 110.60401022845096, 2500, true, 5, 3, 40);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (4, 34, '2026-05-26 11:15:10+07', '2026-05-26 11:15:10+07', 'SPPG Bunut Sanggau', 'Jl. Jenderal Ahmad Yani No.17', '{https://www.facebook.com}', 'Budi Santosa', '08123456789', 'sppg.kapuas.ilirkota@email.com', 0.1235888731544522, 110.60401022845096, 1000, true, 27, 3, 41);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (7, 52, '2026-06-01 11:03:56+07', '2026-06-01 11:03:56+07', 'SPPG Pontianak Timur', 'Jl. Tanjung Raya II No. 88, Pontianak Timur', '{https://www.instagram.com/sppg.pontianaktimur}', 'Siti Rahmawati', '081345678901', 'pontim@sppg.id', -0.0152, 109.3698, 2500, true, 1, 4, 5);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (8, 54, '2026-06-01 11:05:37+07', '2026-06-01 11:05:37+07', 'SPPG Pontianak Timur', 'Jl. Tanjung Raya II No. 88, Pontianak Timur', '{https://www.instagram.com/sppg.pontianaktimur}', 'Siti Rahmawati', '081345678901', 'pontim@sppg.id', -0.0152, 109.3698, 2500, true, 1, 4, 5);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (14, 71, '2026-06-24 20:46:23+07', '2026-06-24 20:46:23+07', NULL, NULL, '{}', NULL, '', '', 0, 0, 0, true, 1, NULL, NULL);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (15, 72, '2026-06-26 20:21:38+07', '2026-06-26 20:21:38+07', 'asas', 'Jl. Ahmad Yani No. 123, Pontianak', '{https://instagram.com/sppg_pontianak,https://facebook.com/sppg.pontianak}', 'Budi Santoso', '081234567890', 'sppg@example.com', -0.02633, 109.342504, 1500, true, 1, 1, 1);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (16, 73, '2026-07-02 20:32:34+07', '2026-07-02 20:32:34+07', 'SPPG Test', NULL, '{}', NULL, '', '', 0, 0, 0, true, 1, NULL, NULL);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (17, 79, '2026-07-02 20:53:47+07', '2026-07-02 20:53:47+07', 'SPPG Test2', NULL, '{}', NULL, '', '', 0, 0, 0, true, 1, NULL, NULL);
INSERT INTO public.sppg (id, user_id, created_at, updated_at, nama, alamat, sosmed_url, kepala_sppg, nomor_telepon, email, latitude, longitude, kapasitas_porsi, status_aktif, version, kecamatan_id, kelurahan_id) VALUES (18, 87, '2026-07-02 22:54:38+07', '2026-07-02 22:54:38+07', 'Test 2', 'Jl. Ahmad Yani No. 123, Pontianak', '{https://instagram.com/sppg_pontianak,https://facebook.com/sppg.pontianak}', 'Budi Santoso', '081234567890', 'sppg@example.com', -0.02633, 109.342504, 1500, false, 1, 1, 1);


ALTER TABLE public.sppg ENABLE TRIGGER ALL;

--
-- Data for Name: alokasi_harian; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.alokasi_harian DISABLE TRIGGER ALL;

INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (1, 5, '2026-06-19', 70000, '2026-06-19 19:19:10.414283', '2026-06-19 19:21:51.775477');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (2, 5, '2026-06-18', 70000, '2026-06-19 19:21:59.60884', '2026-06-19 19:21:59.60884');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (4, 4, '2026-06-20', 9000000, '2026-06-20 18:12:42.534245', '2026-06-20 18:36:59.998524');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (5, 4, '2026-06-21', 9000000, '2026-06-21 10:58:52.40626', '2026-06-21 10:58:52.40626');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (3, 5, '2026-06-20', 80000, '2026-06-20 18:07:49.923017', '2026-06-21 12:59:27.639727');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (6, 5, '2026-06-21', 8700210, '2026-06-21 11:00:20.13874', '2026-06-21 21:31:45.854566');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (7, 4, '2026-06-24', 5000000, '2026-06-24 13:54:06.657179', '2026-06-24 13:54:06.657179');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (8, 5, '2026-07-12', 2999999, '2026-07-12 17:23:37.3622', '2026-07-12 17:23:37.3622');
INSERT INTO public.alokasi_harian (id, sppg_id, tanggal, jumlah, created_at, updated_at) VALUES (9, 5, '2026-07-13', 200000, '2026-07-13 00:12:11.220373', '2026-07-13 00:12:11.220373');


ALTER TABLE public.alokasi_harian ENABLE TRIGGER ALL;

--
-- Data for Name: drivers; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.drivers DISABLE TRIGGER ALL;

INSERT INTO public.drivers (id, created_at, user_id, sppg_id, nama, nomor_telepon, status_aktif, version) VALUES (8, '2026-05-27 18:34:21.052833+07', 47, 4, 'Budi Santoso', '081234567890', true, 1);
INSERT INTO public.drivers (id, created_at, user_id, sppg_id, nama, nomor_telepon, status_aktif, version) VALUES (9, '2026-06-02 16:21:12.754789+07', 55, 4, 'Rian', '08123456789', true, 2);
INSERT INTO public.drivers (id, created_at, user_id, sppg_id, nama, nomor_telepon, status_aktif, version) VALUES (7, '2026-05-27 17:37:41.108617+07', 46, 5, 'Joko Anwar', '082112345978', false, 2);


ALTER TABLE public.drivers ENABLE TRIGGER ALL;

--
-- Data for Name: pedagang_lokal; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.pedagang_lokal DISABLE TRIGGER ALL;

INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (1, 'Toko Sayur Maju', 'Jl. Jenderal Sudirman', '081234567801', 110.598912, 0.124021, 'Sayur', 4, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (2, 'UD Berkah Tani', 'Jl. Pangeran Mas', '081234567802', 110.597635, 0.122584, 'Beras', 5, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (3, 'Segar Abadi', 'Jl. Sutan Syahrir', '081234567803', 110.599041, 0.123867, 'Buah', 7, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (4, 'Sumber Pangan', 'Jl. H. Agus Salim', '081234567804', 110.596882, 0.122963, 'Telur', 8, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (5, 'Toko Makmur', 'Jl. A. Yani', '081234567805', 110.598473, 0.123488, 'Daging Ayam', 14, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (6, 'Mitra Sayuran', 'Jl. Kartini', '081234567806', 110.597944, 0.123952, 'Sayur', 15, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (7, 'Tani Sejahtera', 'Jl. Diponegoro', '081234567807', 110.598281, 0.124172, 'Cabai', 16, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (8, 'Berkah Buah', 'Jl. Gusti Hamzah', '081234567808', 110.598735, 0.122817, 'Buah', 17, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (9, 'Sembako Jaya', 'Jl. RE. Martadinata', '081234567809', 110.597421, 0.123247, 'Sembako', 18, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (10, 'CV Pangan Nusantara', 'Jl. Dr. Wahidin', '081234567810', 110.598642, 0.124426, 'Beras', 4, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (11, 'Toko Rizki', 'Jl. Pangsuma', '081234567811', 110.598986, 0.123015, 'Ikan', 5, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (12, 'Sentra Telur', 'Jl. Seroja', '081234567812', 110.597928, 0.122701, 'Telur', 7, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (13, 'UD Karya Tani', 'Jl. Anggrek', '081234567813', 110.596995, 0.123791, 'Sayur', 8, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (14, 'Fresh Market', 'Jl. Dahlia', '081234567814', 110.599212, 0.124083, 'Buah', 14, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (15, 'Pangan Lestari', 'Jl. Kenanga', '081234567815', 110.597708, 0.123602, 'Beras', 15, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (16, 'Sumber Rezeki', 'Jl. Flamboyan', '081234567816', 110.598152, 0.122925, 'Daging Sapi', 16, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (17, 'Maju Bersama', 'Jl. Cempaka', '081234567817', 110.597551, 0.12431, 'Susu', 17, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (18, 'Panen Segar', 'Jl. Merdeka', '081234567818', 110.599366, 0.123174, 'Sayur', 18, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (19, 'Toko Harapan', 'Jl. Bhayangkara', '081234567819', 110.598561, 0.122536, 'Bawang', 5, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);
INSERT INTO public.pedagang_lokal (id, nama, alamat, no_hp, longitude, latitude, jenis_produk, sppg_id, created_at, updated_at, version) VALUES (20, 'UD Sukses Mandiri', 'Jl. Veteran', '081234567820', 110.598008, 0.123726, 'Minyak Goreng', 7, '2026-07-09 13:02:08.061589+07', '2026-07-09 13:02:08.061589+07', 1);


ALTER TABLE public.pedagang_lokal ENABLE TRIGGER ALL;

--
-- Data for Name: sppg_invitations; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.sppg_invitations DISABLE TRIGGER ALL;

INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (2, 'Zv9DGWJMwZcGTu8B', 'SPPG Bumi Daranante', '2026-07-02 16:39:55.289761+07', NULL, '2026-06-25 16:39:55.290559+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (3, 'uoeqb1oZigQY4U3S', 'SPPG Bukit Bintang', '2026-07-02 16:49:40.918023+07', NULL, '2026-06-25 16:49:40.918704+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (4, 'OVfxFC4bS26Fh-aY', 'SPPG Indonesia', '2026-07-02 18:40:20.260564+07', NULL, '2026-06-25 18:40:20.261764+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (1, 'gevKJcFpfskbZUIe', 'asas', '2026-07-02 16:33:16.231228+07', '2026-06-26 20:21:37.68454+07', '2026-06-25 16:33:16.235819+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (7, '0EOC1yJiU7TP-oRN', 'Test 2', '2026-07-02 22:57:33.447901+07', '2026-07-02 22:54:37.897679+07', '2026-06-25 22:57:33.449933+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (8, 'ZcdFd0K15X0AZVoQ', 'SPPG Test 3', '2026-07-11 21:59:30.111393+07', NULL, '2026-07-04 21:59:30.112465+07');
INSERT INTO public.sppg_invitations (id, token, nama_sppg, expires_at, used_at, created_at) VALUES (9, '_NTNKxLARuuKeDC1', 'SPPG RIAN', '2026-07-11 22:08:12.119402+07', NULL, '2026-07-04 22:08:12.119831+07');


ALTER TABLE public.sppg_invitations ENABLE TRIGGER ALL;

--
-- Data for Name: pending_sppg; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.pending_sppg DISABLE TRIGGER ALL;



ALTER TABLE public.pending_sppg ENABLE TRIGGER ALL;

--
-- Data for Name: pengeluaran_harian; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.pengeluaran_harian DISABLE TRIGGER ALL;

INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (3, 1, 'Beras Premium', 10, 'kg', 18000, '2026-06-20 14:21:29.457244', '2026-06-20 14:21:29.457244', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (6, 5, 'Pupuk', 90, 'kg', 10000, '2026-06-21 10:59:18.345002', '2026-06-21 10:59:18.345002', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (7, 5, 'Beras', 90, 'kg', 9000, '2026-06-21 10:59:59.512289', '2026-06-21 10:59:59.512289', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (9, 1, 'Beras Premium', 10, 'kg', 18000, '2026-06-21 18:10:37.502838', '2026-06-21 18:10:37.502838', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (10, 1, 'Beras Premium', 10, 'kg', 1000, '2026-06-21 18:11:16.619083', '2026-06-21 18:11:16.619083', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (11, 1, 'Beras Premium', 10, 'kg', 1000, '2026-06-21 18:29:59.889349', '2026-06-21 18:29:59.889349', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (12, 1, 'Beras Premium', 10, 'kg', 1000, '2026-06-21 18:35:08.332703', '2026-06-21 18:35:08.332703', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (13, 6, 'Minyak goreng', 20, 'liter', 60000, '2026-06-21 18:53:09.31698', '2026-06-21 18:53:09.31698', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (15, 7, 'Beras', 20, 'Kg', 16000, '2026-06-24 13:54:33.190721', '2026-06-24 13:54:33.190721', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (16, 7, 'Ayam', 40, 'Kg', 40000, '2026-06-24 13:54:52.640872', '2026-06-24 13:54:52.640872', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (17, 7, 'Bayam', 25, 'Kg', 15000, '2026-06-24 13:55:24.924313', '2026-06-24 13:55:24.924313', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (18, 7, 'Susu Segar', 50, 'Liter', 30000, '2026-06-24 13:55:48.341142', '2026-06-24 13:55:48.341142', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (19, 7, 'Pepaya', 40, 'Kg', 8000, '2026-06-24 13:56:14.362097', '2026-06-24 13:56:14.362097', NULL);
INSERT INTO public.pengeluaran_harian (id, alokasi_harian_id, produk, jumlah, satuan, harga_satuan, created_at, updated_at, pedagang_lokal_id) VALUES (20, 7, 'Bumbu', 50, 'Paket', 10000, '2026-06-24 13:57:18.797557', '2026-06-24 13:57:18.797557', NULL);


ALTER TABLE public.pengeluaran_harian ENABLE TRIGGER ALL;

--
-- Data for Name: pengiriman; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.pengiriman DISABLE TRIGGER ALL;

INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (64, '2026-06-06 15:58:53.873968+07', 5, 7, 'sekolah', 52, 'sampai', '2026-06-06 17:11:53.276059+07', '2026-06-06 17:16:09.937732+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (33, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 42, 'dibatalkan', '2026-06-05 11:54:46.186783+07', '2026-06-05 11:55:01.835926+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (174, '2026-07-07 13:22:03.658722+07', 5, 7, 'sekolah', 44, 'sampai', '2026-07-07 13:55:33.75357+07', '2026-07-07 13:56:32.206461+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (34, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 43, 'sampai', '2026-06-05 17:30:45.882457+07', '2026-06-05 17:54:55.421572+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (65, '2026-06-06 15:58:53.873968+07', 5, 7, 'posyandu', 7, 'sampai', '2026-06-06 17:18:33.511226+07', '2026-06-06 17:18:42.58215+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (35, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 44, 'sampai', '2026-06-05 18:02:36.965415+07', '2026-06-05 18:20:36.046347+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (36, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 45, 'sampai', '2026-06-05 18:21:24.947443+07', '2026-06-05 18:21:46.795166+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (66, '2026-06-06 15:58:53.873968+07', 5, 7, 'posyandu', 6, 'sampai', '2026-06-06 17:51:21.573186+07', '2026-06-06 17:51:33.352857+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (37, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 46, 'sampai', '2026-06-05 18:26:52.048437+07', '2026-06-05 18:28:01.698058+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (38, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 47, 'sampai', '2026-06-05 18:29:13.574003+07', '2026-06-05 18:29:52.460009+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (39, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 48, 'sampai', '2026-06-05 18:30:20.206582+07', '2026-06-05 18:32:51.974058+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (71, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 43, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (40, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 49, 'sampai', '2026-06-05 18:33:59.926608+07', '2026-06-05 18:35:21.870538+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (72, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 44, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (41, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 50, 'sampai', '2026-06-05 18:40:34.059889+07', '2026-06-05 18:40:45.304736+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (73, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 45, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (74, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 46, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (75, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 47, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (76, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 48, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (77, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 49, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (78, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 50, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (79, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 51, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (80, '2026-06-06 17:52:13.956956+07', 5, NULL, 'sekolah', 52, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (81, '2026-06-06 17:52:13.956956+07', 5, NULL, 'posyandu', 6, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (47, '2026-06-06 00:03:52.496071+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-06 00:20:08.346027+07', '2026-06-06 00:21:32.342761+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (48, '2026-06-06 00:05:40.670473+07', 5, 7, 'sekolah', 42, 'sampai', '2026-06-06 00:23:17.537096+07', '2026-06-06 00:24:20.901264+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (83, '2026-06-06 17:52:13.956956+07', 5, 7, 'posyandu', 8, 'sampai', '2026-06-06 17:52:32.547031+07', '2026-06-06 17:52:45.314022+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (82, '2026-06-06 17:52:13.956956+07', 5, 7, 'posyandu', 7, 'sampai', '2026-06-06 17:56:30.005516+07', '2026-06-06 18:33:48.000267+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (58, '2026-06-06 00:26:47.827028+07', 5, 7, 'sekolah', 46, 'sampai', '2026-06-06 00:27:00.336758+07', '2026-06-06 00:28:11.709154+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (57, '2026-06-06 00:24:40.611785+07', 5, 7, 'posyandu', 7, 'sampai', '2026-06-06 00:28:23.902201+07', '2026-06-06 00:31:15.869439+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (56, '2026-06-06 00:16:44.10196+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-06 00:31:27.537864+07', '2026-06-06 00:38:44.965645+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (55, '2026-06-06 00:14:42.654933+07', 5, 7, 'sekolah', 51, 'sampai', '2026-06-06 00:39:01.032281+07', '2026-06-06 00:40:46.819808+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (53, '2026-06-06 00:12:12.203806+07', 5, 7, 'sekolah', 46, 'sampai', '2026-06-06 00:41:48.147233+07', '2026-06-06 00:43:43.853747+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (84, '2026-06-06 19:04:56.282064+07', 4, 8, 'sekolah', 32, 'sampai', '2026-06-06 19:05:00.159699+07', '2026-06-06 19:06:19.009802+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (54, '2026-06-06 00:13:37.936617+07', 5, 7, 'sekolah', 43, 'sampai', '2026-06-06 00:46:19.11037+07', '2026-06-06 00:47:00.074406+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (69, '2026-06-06 17:52:13.956956+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-06 19:02:47.18152+07', '2026-06-06 20:15:00.181259+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (52, '2026-06-06 00:11:07.079646+07', 5, 7, 'sekolah', 48, 'sampai', '2026-06-06 00:47:09.687887+07', '2026-06-06 00:47:22.379821+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (30, '2026-06-04 11:16:08.726424+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-04 11:16:56.412153+07', '2026-06-05 10:53:44.056169+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (31, '2026-06-04 11:16:08.726424+07', 5, 7, 'sekolah', 42, 'sampai', '2026-06-05 11:37:45.597778+07', '2026-06-05 11:44:03.688233+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (50, '2026-06-06 00:10:06.775265+07', 5, 7, 'posyandu', 7, 'sampai', '2026-06-06 00:47:35.093656+07', '2026-06-06 00:47:43.1686+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (85, '2026-06-06 19:04:56.282064+07', 4, 8, 'sekolah', 37, 'sampai', '2026-06-06 21:12:06.735496+07', '2026-06-06 21:41:08.769936+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (88, '2026-06-06 22:03:38.675719+07', 4, NULL, 'posyandu', 9, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (86, '2026-06-06 22:03:38.675719+07', 4, 8, 'posyandu', 3, 'sampai', '2026-06-06 22:03:54.159372+07', '2026-06-06 22:15:11.189391+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (87, '2026-06-06 22:03:38.675719+07', 4, 8, 'posyandu', 4, 'sampai', '2026-06-06 22:16:02.587931+07', '2026-06-07 00:32:56.236791+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (32, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-05 11:51:34.072697+07', '2026-06-05 11:51:47.694471+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (49, '2026-06-06 00:08:32.726352+07', 5, 7, 'sekolah', 49, 'dibatalkan', '2026-06-06 00:47:50.255965+07', '2026-06-06 00:47:54.759142+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (91, '2026-06-07 08:44:37.491421+07', 4, NULL, 'sekolah', 38, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (92, '2026-06-07 08:44:37.491421+07', 4, NULL, 'sekolah', 53, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (59, '2026-06-06 00:49:01.193639+07', 5, 7, 'sekolah', 45, 'sampai', '2026-06-06 00:49:14.202968+07', '2026-06-06 00:49:19.418346+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (93, '2026-06-07 08:44:37.491421+07', 4, NULL, 'sekolah', 57, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (51, '2026-06-06 00:10:38.125527+07', 5, 7, 'sekolah', 51, 'sampai', '2026-06-06 15:44:44.715089+07', '2026-06-06 15:46:56.412144+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (94, '2026-06-07 08:44:37.491421+07', 4, NULL, 'sekolah', 58, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (60, '2026-06-06 00:49:01.193639+07', 5, 7, 'sekolah', 47, 'sampai', '2026-06-06 15:47:15.97439+07', '2026-06-06 15:49:37.027834+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (95, '2026-06-07 08:44:37.491421+07', 4, NULL, 'sekolah', 59, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (96, '2026-06-07 08:44:37.491421+07', 4, NULL, 'posyandu', 3, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (97, '2026-06-07 08:44:37.491421+07', 4, NULL, 'posyandu', 4, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (61, '2026-06-06 15:49:51.26863+07', 5, 7, 'sekolah', 52, 'sampai', '2026-06-06 15:50:04.432014+07', '2026-06-06 15:56:40.0612+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (98, '2026-06-07 08:44:37.491421+07', 4, NULL, 'posyandu', 9, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (62, '2026-06-06 15:49:51.26863+07', 5, 7, 'posyandu', 6, 'sampai', '2026-06-06 15:56:56.891821+07', '2026-06-06 15:57:18.3618+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (70, '2026-06-06 17:52:13.956956+07', 5, 7, 'sekolah', 42, 'sampai', '2026-06-06 20:43:10.676329+07', '2026-06-07 09:08:20.397787+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (67, '2026-06-06 15:58:53.873968+07', 5, 7, 'sekolah', 50, 'sampai', '2026-06-06 16:01:31.382248+07', '2026-06-06 16:07:00.061164+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (68, '2026-06-06 15:58:53.873968+07', 5, 7, 'sekolah', 51, 'sampai', '2026-06-06 16:07:15.697249+07', '2026-06-06 16:07:39.829722+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (63, '2026-06-06 15:57:06.867028+07', 5, 7, 'posyandu', 6, 'sampai', '2026-06-06 16:07:56.671239+07', '2026-06-06 16:10:28.474738+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (105, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 47, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (106, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 48, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (107, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 49, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (100, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 42, 'sampai', '2026-06-07 09:17:11.399689+07', '2026-06-07 09:17:39.541732+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (102, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 44, 'sampai', '2026-06-07 09:25:58.458566+07', '2026-06-07 09:41:43.571548+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (89, '2026-06-07 08:44:37.491421+07', 4, 8, 'sekolah', 32, 'sampai', '2026-06-07 08:45:01.074518+07', '2026-06-07 10:18:48.655483+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (103, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 45, 'sampai', '2026-06-07 09:57:50.892743+07', '2026-06-07 10:21:13.564995+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (90, '2026-06-07 08:44:37.491421+07', 4, 8, 'sekolah', 37, 'sampai', '2026-06-07 14:04:33.985628+07', '2026-06-07 14:44:17.814817+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (42, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 51, 'sampai', '2026-06-09 09:14:47.992329+07', '2026-06-09 09:51:29.632127+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (44, '2026-06-05 11:51:22.506847+07', 5, 7, 'posyandu', 6, 'sampai', '2026-06-09 09:58:27.74499+07', '2026-06-09 10:01:18.909098+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (43, '2026-06-05 11:51:22.506847+07', 5, 7, 'sekolah', 52, 'sampai', '2026-06-09 10:01:24.340118+07', '2026-06-09 10:23:18.545949+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (45, '2026-06-05 11:51:22.506847+07', 5, 7, 'posyandu', 7, 'sampai', '2026-06-09 10:23:25.715227+07', '2026-06-09 10:23:45.461366+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (108, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 50, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (109, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 51, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (110, '2026-06-07 09:09:59.933968+07', 5, NULL, 'sekolah', 52, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (111, '2026-06-07 09:09:59.933968+07', 5, NULL, 'posyandu', 6, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (112, '2026-06-07 09:09:59.933968+07', 5, NULL, 'posyandu', 7, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (113, '2026-06-07 09:09:59.933968+07', 5, NULL, 'posyandu', 8, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (99, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 41, 'sampai', '2026-06-07 09:10:28.763081+07', '2026-06-07 09:15:23.028089+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (175, '2026-07-07 13:22:03.658722+07', 5, 7, 'sekolah', 45, 'berangkat', '2026-07-07 13:57:50.347724+07', NULL, 2);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (101, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 43, 'sampai', '2026-06-07 09:24:19.401986+07', '2026-06-07 09:24:51.022488+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (104, '2026-06-07 09:09:59.933968+07', 5, 7, 'sekolah', 46, 'sampai', '2026-06-07 14:44:46.905752+07', '2026-06-08 22:55:49.502471+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (46, '2026-06-05 11:51:22.506847+07', 5, 7, 'posyandu', 8, 'sampai', '2026-06-09 10:23:48.198928+07', '2026-06-12 19:31:04.023889+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (114, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 41, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (115, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 42, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (116, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 43, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (117, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 45, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (118, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 44, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (119, '2026-06-12 19:46:00.177947+07', 5, NULL, 'sekolah', 46, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (122, '2026-06-12 20:11:16.490232+07', 5, NULL, 'sekolah', 49, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (123, '2026-06-12 20:18:10.047557+07', 5, NULL, 'sekolah', 51, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (124, '2026-06-12 20:19:42.743121+07', 5, NULL, 'sekolah', 52, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (125, '2026-06-12 20:22:44.384627+07', 5, NULL, 'posyandu', 6, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (126, '2026-06-12 20:24:21.820355+07', 5, NULL, 'posyandu', 7, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (127, '2026-06-12 20:25:55.900272+07', 5, NULL, 'sekolah', 47, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (128, '2026-06-12 20:26:57.379901+07', 5, NULL, 'sekolah', 44, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (129, '2026-06-12 20:29:42.198217+07', 5, NULL, 'sekolah', 48, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (120, '2026-06-12 19:46:00.177947+07', 5, 7, 'sekolah', 47, 'sampai', '2026-06-12 21:29:50.553766+07', '2026-06-12 21:54:23.160357+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (121, '2026-06-12 19:46:00.177947+07', 5, 7, 'sekolah', 48, 'sampai', '2026-06-12 22:06:18.622334+07', '2026-06-13 12:02:13.118958+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (130, '2026-06-13 12:46:12.398122+07', 5, NULL, 'sekolah', 41, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (131, '2026-06-13 12:46:12.398122+07', 5, NULL, 'sekolah', 42, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (132, '2026-06-13 12:46:12.398122+07', 5, NULL, 'sekolah', 43, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (133, '2026-06-13 12:49:03.577413+07', 5, 7, 'sekolah', 51, 'sampai', '2026-06-13 12:49:50.48574+07', '2026-06-13 13:17:37.884218+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (134, '2026-06-13 13:04:33.143529+07', 5, 7, 'posyandu', 6, 'sampai', '2026-06-13 13:39:27.348809+07', '2026-06-13 13:44:01.516056+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (135, '2026-06-14 20:53:37.239624+07', 4, NULL, 'sekolah', 32, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (136, '2026-06-14 20:53:37.239624+07', 4, NULL, 'sekolah', 37, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (137, '2026-06-14 20:53:37.239624+07', 4, NULL, 'sekolah', 38, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (138, '2026-06-14 20:53:37.239624+07', 4, NULL, 'sekolah', 53, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (139, '2026-06-14 22:21:08.441937+07', 5, NULL, 'sekolah', 41, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (141, '2026-06-14 22:21:08.441937+07', 5, NULL, 'sekolah', 43, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (143, '2026-06-14 22:21:08.441937+07', 5, NULL, 'sekolah', 45, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (144, '2026-06-14 22:21:08.441937+07', 5, NULL, 'sekolah', 46, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (140, '2026-06-14 22:21:08.441937+07', 5, 7, 'sekolah', 42, 'sampai', '2026-06-14 22:24:39.082755+07', '2026-06-14 22:26:48.6102+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (142, '2026-06-14 22:21:08.441937+07', 5, 7, 'sekolah', 44, 'sampai', '2026-06-14 22:35:21.35944+07', '2026-06-14 22:36:18.197441+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (145, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 32, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (146, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 37, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (147, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 38, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (148, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 53, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (149, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 57, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (150, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 58, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (151, '2026-06-19 12:21:11.05824+07', 4, NULL, 'sekolah', 59, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (152, '2026-06-19 12:21:11.05824+07', 4, NULL, 'posyandu', 3, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (153, '2026-06-19 12:21:11.05824+07', 4, NULL, 'posyandu', 4, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (154, '2026-06-19 12:21:11.05824+07', 4, NULL, 'posyandu', 9, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (155, '2026-06-24 13:57:51.976047+07', 4, 8, 'sekolah', 32, 'sampai', '2026-06-24 13:58:24.797466+07', '2026-06-24 13:59:35.210766+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (156, '2026-06-24 13:57:51.976047+07', 4, 8, 'sekolah', 37, 'sampai', '2026-06-24 13:59:41.854309+07', '2026-06-24 13:59:50.276848+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (157, '2026-06-24 13:57:51.976047+07', 4, 8, 'sekolah', 38, 'sampai', '2026-06-24 14:00:59.04686+07', '2026-06-24 14:01:51.531948+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (158, '2026-06-24 13:57:51.976047+07', 4, 8, 'sekolah', 53, 'sampai', '2026-06-24 14:01:56.136282+07', '2026-06-24 14:08:07.816622+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (159, '2026-06-24 14:10:45.109862+07', 4, 8, 'posyandu', 4, 'sampai', '2026-06-24 14:11:56.879728+07', '2026-06-24 14:16:15.430834+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (162, '2026-07-05 22:17:49.969191+07', 4, NULL, 'sekolah', 38, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (163, '2026-07-05 22:17:49.969191+07', 4, NULL, 'sekolah', 53, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (164, '2026-07-05 22:17:49.969191+07', 4, NULL, 'sekolah', 57, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (165, '2026-07-05 22:17:49.969191+07', 4, NULL, 'sekolah', 58, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (166, '2026-07-05 22:17:49.969191+07', 4, NULL, 'sekolah', 59, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (167, '2026-07-05 22:17:49.969191+07', 4, NULL, 'posyandu', 3, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (168, '2026-07-05 22:17:49.969191+07', 4, NULL, 'posyandu', 4, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (169, '2026-07-05 22:17:49.969191+07', 4, NULL, 'posyandu', 9, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (160, '2026-07-05 22:17:49.969191+07', 4, 8, 'sekolah', 32, 'sampai', '2026-07-05 22:23:06.391044+07', '2026-07-05 22:38:55.833912+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (170, '2026-07-05 22:42:27.668052+07', 4, NULL, 'sekolah', 32, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (161, '2026-07-05 22:17:49.969191+07', 4, 8, 'sekolah', 37, 'sampai', '2026-07-05 22:39:10.147732+07', '2026-07-05 22:46:07.633142+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (176, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 46, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (177, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 47, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (178, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 48, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (179, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 49, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (180, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 50, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (181, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 51, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (182, '2026-07-07 13:22:03.658722+07', 5, NULL, 'sekolah', 52, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (183, '2026-07-07 13:22:03.658722+07', 5, NULL, 'posyandu', 6, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (184, '2026-07-07 13:22:03.658722+07', 5, NULL, 'posyandu', 7, 'menunggu', NULL, NULL, 1);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (171, '2026-07-07 13:22:03.658722+07', 5, 7, 'sekolah', 41, 'sampai', '2026-07-07 13:22:14.01007+07', '2026-07-07 13:24:53.919166+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (172, '2026-07-07 13:22:03.658722+07', 5, 7, 'sekolah', 42, 'sampai', '2026-07-07 13:29:52.927269+07', '2026-07-07 13:42:24.876372+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (173, '2026-07-07 13:22:03.658722+07', 5, 7, 'sekolah', 43, 'sampai', '2026-07-07 13:52:04.357582+07', '2026-07-07 13:54:31.785441+07', 3);
INSERT INTO public.pengiriman (id, created_at, sppg_id, driver_id, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, version) VALUES (185, '2026-07-07 13:22:03.658722+07', 5, NULL, 'posyandu', 8, 'menunggu', NULL, NULL, 1);


ALTER TABLE public.pengiriman ENABLE TRIGGER ALL;

--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.permissions DISABLE TRIGGER ALL;

INSERT INTO public.permissions (id, code) VALUES (1, 'sekolah:read');
INSERT INTO public.permissions (id, code) VALUES (2, 'sekolah:write');


ALTER TABLE public.permissions ENABLE TRIGGER ALL;

--
-- Data for Name: posyandu; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.posyandu DISABLE TRIGGER ALL;

INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (6, '2026-05-26 21:10:20+07', '2026-05-26 21:10:20+07', 5, 'Posyandu Melati Ilir Kota', 'Jl. Melati Ilir Kota', 0.1237, 110.6041, 95, 22, 1, 3, 40);
INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (7, '2026-05-26 21:10:20+07', '2026-05-26 21:10:20+07', 5, 'Posyandu Mawar Ilir Kota', 'Jl. Mawar Ilir Kota', 0.1234, 110.6037, 88, 19, 1, 3, 40);
INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (8, '2026-05-26 21:10:20+07', '2026-05-26 21:10:20+07', 5, 'Posyandu Anggrek Ilir Kota', 'Jl. Anggrek Ilir Kota', 0.124, 110.6046, 103, 24, 1, 3, 40);
INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (3, '2026-05-26 21:09:20+07', '2026-05-26 21:09:20+07', 4, 'Posyandu Bersama IIS', 'Jl. Melati Bunutm', 0.1268353398297325, 110.5690282709128, 85, 18, 8, 3, 164);
INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (4, '2026-05-26 21:09:20+07', '2026-05-26 21:09:20+07', 4, 'Posyandu Mawar Bunut', 'Jl. Mawar Bunut', 0.1215, 110.5806, 92, 21, 3, 3, 164);
INSERT INTO public.posyandu (id, created_at, updated_at, sppg_id, nama, alamat, latitude, longitude, jumlah_balita, jumlah_ibu_hamil, version, kecamatan_id, kelurahan_id) VALUES (9, '2026-06-01 17:29:19+07', '2026-06-01 17:29:19+07', 4, 'Posyandu Mawar', 'Jl. Adisucipto Km. 10', 0.1268353398297325, 110.5690282709128, 120, 24, 3, 3, 164);


ALTER TABLE public.posyandu ENABLE TRIGGER ALL;

--
-- Data for Name: produksi_harian; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.produksi_harian DISABLE TRIGGER ALL;

INSERT INTO public.produksi_harian (id, sppg_id, tanggal, waktu_mulai, estimasi_waktu_selesai, created_at, updated_at) VALUES (1, 5, '2026-06-21', '2026-06-21 04:30:00', '2026-06-21 16:50:00', '2026-06-21 01:45:06.052462', '2026-06-21 21:43:40.328191');
INSERT INTO public.produksi_harian (id, sppg_id, tanggal, waktu_mulai, estimasi_waktu_selesai, created_at, updated_at) VALUES (2, 4, '2026-06-24', '2026-06-24 04:00:00', '2026-06-24 10:00:00', '2026-06-24 13:53:49.193408', '2026-06-24 13:53:49.193408');
INSERT INTO public.produksi_harian (id, sppg_id, tanggal, waktu_mulai, estimasi_waktu_selesai, created_at, updated_at) VALUES (3, 5, '2026-07-09', '2026-07-09 00:00:00', '2026-07-09 20:00:00', '2026-07-09 12:38:51.488063', '2026-07-09 12:38:51.488063');


ALTER TABLE public.produksi_harian ENABLE TRIGGER ALL;

--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.schema_migrations DISABLE TRIGGER ALL;

INSERT INTO public.schema_migrations (version, dirty) VALUES (34, false);


ALTER TABLE public.schema_migrations ENABLE TRIGGER ALL;

--
-- Data for Name: sekolah; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.sekolah DISABLE TRIGGER ALL;

INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (32, '2026-05-26 21:08:41+07', 'SD Negeri 04 Bunut', 'SD', 420, 0.1268353398297325, 110.5690282709128, 'Jl. Bunut 4', 3, 4, '2026-05-26 21:08:41+07', 3, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (37, '2026-05-26 21:08:41+07', 'SMP Negeri 04 Bunut', 'SMP', 720, 0.1268353398297325, 110.5690282709128, 'Jl. SMP Bunut 4', 3, 4, '2026-05-26 21:08:41+07', 3, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (38, '2026-05-26 21:08:41+07', 'SMA Negeri 01 Bunut', 'SMA', 920, 0.1268353398297325, 110.5690282709128, 'Jl. SMA Bunut 1', 3, 4, '2026-05-26 21:08:41+07', 3, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (41, '2026-05-26 21:09:49+07', 'SD Negeri 01 Ilir Kota', 'SD', 390, 0.12770046445395156, 110.60128011718128, 'Jl. Ilir Kota 1', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (42, '2026-05-26 21:09:49+07', 'SD Negeri 02 Ilir Kota', 'SD', 410, 0.12770046445395156, 110.60128011718128, 'Jl. Ilir Kota 2', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (43, '2026-05-26 21:09:49+07', 'SD Negeri 03 Ilir Kota', 'SD', 430, 0.12770046445395156, 110.60128011718128, 'Jl. Ilir Kota 3', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (44, '2026-05-26 21:09:49+07', 'SD Negeri 04 Ilir Kota', 'SD', 370, 0.12770046445395156, 110.60128011718128, 'Jl. Ilir Kota 4', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (45, '2026-05-26 21:09:49+07', 'SD Negeri 05 Ilir Kota', 'SD', 400, 0.12770046445395156, 110.60128011718128, 'Jl. Ilir Kota 5', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (46, '2026-05-26 21:09:49+07', 'SMP Negeri 01 Ilir Kota', 'SMP', 640, 0.12770046445395156, 110.60128011718128, 'Jl. SMP Ilir Kota 1', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (47, '2026-05-26 21:09:49+07', 'SMP Negeri 02 Ilir Kota', 'SMP', 690, 0.12770046445395156, 110.60128011718128, 'Jl. SMP Ilir Kota 2', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (48, '2026-05-26 21:09:49+07', 'SMP Negeri 03 Ilir Kota', 'SMP', 720, 0.12770046445395156, 110.60128011718128, 'Jl. SMP Ilir Kota 3', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (49, '2026-05-26 21:09:49+07', 'SMP Negeri 04 Ilir Kota', 'SMP', 670, 0.12770046445395156, 110.60128011718128, 'Jl. SMP Ilir Kota 4', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (50, '2026-05-26 21:09:49+07', 'SMA Negeri 01 Ilir Kota', 'SMA', 910, 0.12770046445395156, 110.60128011718128, 'Jl. SMA Ilir Kota 1', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (51, '2026-05-26 21:09:49+07', 'SMA Negeri 02 Ilir Kota', 'SMA', 940, 0.12770046445395156, 110.60128011718128, 'Jl. SMA Ilir Kota 2', 1, 5, '2026-05-26 21:09:49+07', 3, 40);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (53, '2026-05-31 00:44:15+07', 'SMA Negeri 50 Sanggau', 'SMA', 950, 0.1268353398297325, 110.5690282709128, 'Jl. Ahmad Yani No. 90', 3, 4, '2026-05-31 00:44:15+07', 3, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (57, '2026-06-01 23:47:06+07', 'asas', 'SD', 33, 0.1268353398297325, 110.5690282709128, 'asasa', 2, 4, '2026-06-01 23:47:06+07', 5, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (58, '2026-06-02 00:00:12+07', 'asas', 'SD', 888, 0.1268353398297325, 110.5690282709128, 'asassa', 2, 4, '2026-06-02 00:00:12+07', 5, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (59, '2026-06-02 00:01:52+07', 'BBB', 'SD', 3, 0.1268353398297325, 110.5690282709128, 'asas', 2, 4, '2026-06-02 00:01:52+07', 5, 164);
INSERT INTO public.sekolah (id, created_at, nama, tingkat, jumlah_siswa, latitude, longitude, alamat, version, sppg_id, updated_at, kecamatan_id, kelurahan_id) VALUES (52, '2026-05-26 21:09:49+07', 'SMA Negeri 03 Ilir Kota', 'SMA', 890, 0.12770046445395156, 110.60128011718128, 'Jl. SMA Ilir Kota 3', 1, 5, '2026-05-26 21:09:49+07', 3, 40);


ALTER TABLE public.sekolah ENABLE TRIGGER ALL;

--
-- Data for Name: tokens; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.tokens DISABLE TRIGGER ALL;

INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xa1e0969152d79cfb00adcc941cbcc77193460157408005b873bdf84deda67b01', 34, '2026-06-15 20:53:23+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xcf993408774a91aa4c4478b0f7f875d49640396cfa2ad533f34c8f31c7bc94c3', 46, '2026-06-19 13:37:54+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf936103de68fb07de615a8cdfa3ac7dddafbf9c0aeab809c9b10732708a9d624', 34, '2026-06-19 13:38:41+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd1f6462471a9251136f44488f91a5e78e41ab2da8e42759b281696f554037337', 46, '2026-06-19 14:29:41+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9fda2e6059eacfe3c00b96bb35c779cb69e6ce085339fdd678bf1221bdfb2628', 37, '2026-06-19 14:29:54+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x26aad3564eebe838851555937764477e1cc0fed16cad8314b76959e58fe1d1db', 46, '2026-06-19 14:30:46+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5b976de886b06b16d66e9d83f280f04b19221320dcd9515871b519d558b4955e', 46, '2026-06-20 19:04:40+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x077628652c3c133946cf752fc417b797ac8216c7ea56bcf4f4970911cc9e481f', 39, '2026-06-20 19:06:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x3f32ad17ad49882a220c8e7bc2ca51b829f6c3d9a2a691a5e88b27561c9e025e', 46, '2026-06-20 21:56:16+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x001988b43197a84411ceecaea7884b68bc6d35b60918220e77a3bc580163d02d', 46, '2026-06-20 22:08:55+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf5228b112a609d4fbe9671f455787d0d48ebf3fdd75b8f1db6cce0152d2b9e56', 34, '2026-06-21 17:51:53+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xab60c7ca064c004d8cfcb897f03d1931016cb3eb8a1e3369047d8bda2123e12d', 39, '2026-06-01 20:58:34+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x29a3cad66f6149db5f7c4c649d63313c9719818ae12a291da5fd879c471ac0fc', 39, '2026-06-21 19:47:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf4c7478e0e3ce5c301ad63def2968a3461bff8f31cbe57908e11849b5fd3d81b', 34, '2026-06-21 23:36:18+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf0d5a8b95015b2246ddc2bc00174f8ac5518bb2d005d66b4bad1fc4eed50b8d8', 34, '2026-06-21 23:58:40+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5bafded432ab377aadfa80705d52bc71779da80691e4bd5e85b23e922ba2746d', 46, '2026-06-22 00:26:00+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xea7710df0656ea28e0adfe456cfc4db14ac6666a26862aad453a7a717c8f9942', 39, '2026-06-22 11:00:11+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x20fa82dac79722bd3de5d02f41ab48842cf5997328ff6bff6562a0020f426235', 34, '2026-06-24 17:21:30+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xc5f83f97067c71efc095ae83698da5c31a925b055281917a7fb366fdf86c3923', 37, '2026-06-25 10:36:55+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xde9ae87bc80d3e1eaaecf708328951f3816eb28f2e0c0a9f7f129ecb5945363f', 37, '2026-06-25 13:13:12+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xcc9b9afca5e3d7c4e9f6d3455736da053b1eba65cece4d003c159969027083f7', 34, '2026-06-25 13:42:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x40848f1873bddd5d54a1dd5023ef7b8052a8186b93df56ae97d4aa77cf7f931f', 46, '2026-06-25 13:58:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x45030a38308002a4691634c78f2ae2eeb9c49990f44b74a2a70826e84f2d5b56', 47, '2026-06-25 13:58:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x47528023bd9e1285a9ea4b62bd924cb52fbfbb6d1bf527e17e43e0611651257c', 47, '2026-06-25 14:10:56+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xbe9819b1c871b9e6da86d4f696e2ede7727ba193554aaf4df2ae9f23acc6edc0', 34, '2026-06-25 14:26:37+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd3997159a772a552bb6fe848651bdf05dfc8329f76c8738a62eea5e13b7a5d92', 37, '2026-06-25 14:27:00+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x53444f9ef16de731a5fae500e3bbd12c1a3c8f441e8b4816b54756518e799a71', 37, '2026-06-26 00:58:12+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x819f2e43a39f9b023553534972611da106279bd2e542ac49532bc8b12be62cb1', 37, '2026-06-27 23:44:13+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xa094a1f00c3a70e0fa32dd5da6611edae7fa6eaf1e5eb099eb123350813f4b31', 87, '2026-07-04 10:00:37+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf0ef3ba14fe134122a216a998e14af488e47c163b3c0752fcc87872a9554fd1e', 37, '2026-07-05 01:21:28+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4f8e8cd288df4a9b40c55190716fc07c29eeb0b3ec2404c1a2775be08a711b0d', 48, '2026-06-04 09:22:42+07', 'activation');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x78928e71055e009458daf0a5d001a95f36c6724b1541488148addcde2027ae9e', 37, '2026-06-02 10:44:53+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2e1b963cfb5ed9cbb96b4d1a835809d9cfd14eb0df95f612cd632ecda6351a05', 34, '2026-06-02 11:38:02+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb8485a7a0947246f2b612999418a1b949155da86c3ed0c060a30dcaebb0eabee', 34, '2026-06-02 12:41:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x7ef90f7536cfb8bbb99fd615152e6b8278c5ec3d989af732540f12da331b2bf8', 34, '2026-06-02 12:41:26+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x26c20af85669325c6d1e5d8e9d963fed26df5ef76277468571e05c6ca1462759', 34, '2026-06-03 13:29:51+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1920e571ecd29109233cd3404c60dac87ea642a71253de99042585235710afe2', 39, '2026-06-03 18:38:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x665fb135fbd6b186ce07a1c157bfeaea2f95bdac3445bbb447ebac089a449f78', 34, '2026-06-03 18:39:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x712e040b3b8d2c9290792978a30fe100ca5188cb246b8bbe680ad88de6cc3377', 39, '2026-06-03 18:39:56+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xefbb5e47929ad8047bfc732b3e12d30b5d9fc099c0aea92b298d353a8b77c96e', 34, '2026-06-04 09:59:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x329c6e72c4c65b509f1845d95b2ed8a77e208f0948fcb24fe2698260fb4c5e39', 39, '2026-06-04 10:01:12+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x03fcff1b12f9fdbc281162fece1be5a945e4b3d058efa5fc487f928caf039e59', 39, '2026-06-04 18:48:11+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xdc230d0537820bc3feaeec5ed5e58cd3f47df994d5c3ce5cf38f2869b7a1e116', 39, '2026-06-04 20:30:01+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1b85e23c297fbfa096612053b54d4eaee626b3abd766c0b02e96fb78135885bb', 34, '2026-06-04 20:34:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x6616af99a5b1f5fbc29a05bcaa6ab9a1cd20b0a5bb41cb638324c963d3a8fccb', 39, '2026-06-04 20:35:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xa906cf026b30b4b943a674cddd8a098b9a759bf31fb36027c53538672e85a2ad', 34, '2026-06-04 21:19:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9be54991264620393c8f76e24617bf5726e1674bf9e3eaac45e23bab85489dea', 47, '2026-06-04 22:43:37+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x26189b8ac51cca9af8fe338178fa641e5da03166e95f8a81c4b654b04d6b07cf', 34, '2026-06-05 02:08:46+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xc212aff3bff8175280709587a86f6eef7154b07c13d5bd27c82e2e27d8c7021f', 39, '2026-06-05 02:09:02+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb23d757995e3467300d20f395794cf84bbb130f828d2d90a70f76235c77d455d', 34, '2026-06-05 10:55:46+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x293f242abaa199583fd6d28ce0f8789db5fc9b9f11741952b412e7031a3705a5', 34, '2026-06-05 10:56:52+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xad56a9551df20bf5d22899e8a675b11edc2cccf0e688e4e6bb0908707271ee87', 39, '2026-06-05 10:57:29+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2b145ff51dae590203fd3edd04de2ff5cdcbb6c947e5882d56e0e4cf880161d4', 34, '2026-06-05 11:10:35+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x797af55490ac2075e617b88e0dc5db53782f466ea49b7009b7980fc5f511f9e2', 39, '2026-06-05 11:12:09+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x3dfefdf469e09ad9aa19c9f46372a3b242af0ad4ff0edabdaeac2d14607f4c3c', 34, '2026-06-05 11:12:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x81fb0a0298408a4eee5dc97de56f6a8524905ddc195d3f398ed9664604a40ecc', 34, '2026-06-05 11:14:48+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x19dd5a98ddb21cbadfa8c99389ae359634edcaf4d935f7149b4fb28fd3807477', 39, '2026-06-05 11:14:59+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd76abd957f8ff1e4b435f7766c68636110810ba672e8f6d6cb5488fcab29417d', 34, '2026-06-06 11:50:19+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x230518601c9c0f96be4ae6913d51f2574a198a3008940a9caa10b8dad78fc337', 39, '2026-06-06 11:51:01+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfd049dc10bc817a98141d39f3bee9b2f4069ba59838ac85efc6ee20865e32ebc', 34, '2026-06-06 17:29:15+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x11175729a79e7dd3c6ba23e37c6610551d6a0801944975d2df85f44ab3cbf56f', 39, '2026-06-06 17:30:26+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd46efa172780e29a059d9f46ee60c00d7c2b2d7fd7c4a4e3e11589e3d0a6ebad', 34, '2026-06-06 23:52:58+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xe220b6cc9f8d98c02ec05c30261ad16440a6b870abbf2a0da60235ecdabacd4c', 39, '2026-06-06 23:59:15+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xbca623a205994ada917a202a7b6d8534a6bf72ebdd56d071506b580c3c6bd8f3', 39, '2026-06-07 00:13:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x024d80790840a39300811b0c2be1ca3ed845055ce1bee0dfde40fbee22a42a12', 39, '2026-06-07 00:23:04+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfc2d7e8fc1e1a8f19f1c460f7804dea3b8e098a1630f5e27ec29d7ad3633e15f', 39, '2026-06-07 09:24:40+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x7debb1823466dd14c2111b46dec1a03b6884dc9b205254f94ad0e069cb033c28', 39, '2026-06-07 10:16:31+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1de22be5fc6e7c151c7054279e14f2587a2c2080bf876c3b04f0776032bc2308', 34, '2026-06-07 12:02:25+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x99a420e6bd3fcd2750e7f2537cb40f8be89357d48f450382c65565e0a7ddfb80', 39, '2026-06-07 15:42:25+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x441e131e4ccf589be94d9357510efd07f158f698ad585f3c3081408ec47893c2', 39, '2026-06-07 15:59:56+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x19180ebc90a5f97caf457fbca0a959f725b6056f85a9c7b85731ac3549ec477f', 47, '2026-06-07 19:04:17+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5147f26bb41952ad561b307d059ff36cdcc51c1d90c33b0e2bcaeba0fe3f2144', 34, '2026-06-07 19:04:44+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5afd97277be92b7d653f1369b2cf83655e7097a1f03c74904dd96f6c9b00f7b1', 47, '2026-06-07 20:04:07+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xdedcc0ddb80e8968f99d4e4e3293b9f63e917f83d3b08c3289a137384636b572', 47, '2026-06-07 21:12:02+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x6b5551ff8dfe4537517fd69b67f0f1967248aec560126fa1417da983c18d06b8', 47, '2026-06-07 21:38:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x233acda193dea7064b47838e038009beedd19472969e9f4dcddbb47c9be2d481', 34, '2026-06-07 22:03:23+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x0d423227186deafb70677dcb3adc1b1541fc9429d688d294804cdaf32c286571', 34, '2026-06-07 23:30:48+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2d8a890789ada7ba3f2e67d4243dae5727ddedddb053270d7864bde32a024e2d', 39, '2026-06-08 09:09:36+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xbdb9265223153d70365e0d5ee161e3d72dd729cec41563b125d318fe6e1f1fe8', 34, '2026-06-08 09:42:04+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x0a0d118ea953fdda807ffd93971ff807bb870160923f9f4e3aad7e880dc01a0b', 47, '2026-06-08 09:57:04+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2980f7363797a86c2a1232abc962566d835a5bb8d252ae1368ac2114e11e653f', 46, '2026-06-15 20:42:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xe3557ff0c5cf42e7e1cda67daf26f70571faa8665af24598da3fd3708a888393', 47, '2026-06-09 16:55:01+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfcdc6a5265f9182c1dfdf124e756f75cc5fa5479c9853114860dacf3029887a3', 46, '2026-06-15 20:50:29+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x790302349113c79fb76c9ff029bd5abd4ee171ed1dffebd2b6cff7fa759ea8d7', 39, '2026-06-15 20:56:39+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x72ee5f31e50e79ce5317eca048e164d47ecedd676a61ee83736d8206923f72f5', 37, '2026-06-11 15:55:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9e085449ca6e5d906a6956c6512dff2047fb41d54247d0ee13eb67d4957005bb', 34, '2026-06-11 15:56:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfa56755c2c7af32fd0f5481eeb8dffe1ac93f4be5d58a29769aae7600b47c332', 48, '2026-06-11 15:57:00+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4138114009221c819b46b9c95b7eccb709796ca3790745889564efbaca39bd13', 46, '2026-06-15 21:02:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x540c1d994687ef9546ae4a3f333688af322158df93c8dcc5e07d65bf1d0c73f9', 37, '2026-06-11 15:57:14+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x6a90106a8c0fa47ded6db444c55e6256b77b3bd089b60e1df26f822d93af3140', 34, '2026-06-11 15:58:33+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x02ddb9fe17788a557868d3b2d69e5dbb7ed6b82b07b68ceea19f6a8a25836f4a', 39, '2026-06-11 16:49:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9915fc43023638ea7347feec43c83c2450cd60ca2df90abec6a8880d3bfaf9be', 47, '2026-06-11 16:54:36+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xeff2a4077c39e45cb8bb7a416ed6863b8505b6ba1034a616d8868a526efef425', 46, '2026-06-15 21:08:14+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xc2f0275099f64907e389719a1df64ee66112be7f1fffee8efcb01da6bb0456f6', 46, '2026-06-15 22:31:55+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x92055ab51de6c97c60f6dd17ec4f37411b300bc94bfcc4a27ab0f782b673e20b', 39, '2026-06-12 17:48:27+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xde15327aca5caf8c95de0e9f4d5b8a6b8f0e1515785e745481426fa8fb7b7ff8', 39, '2026-06-13 19:45:39+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9506d233a1025a8cadea6cd7654f3a09028378d06032d10829ea51631bd5ebd1', 47, '2026-06-19 13:36:44+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9a063d490ac97e89b0a9b1dcc4e592594bdadaaa987675b668981865c6a04c05', 34, '2026-06-19 13:37:15+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xe00cd0f44c23212593887cda9f4116e39654ec3531176e6ea01d73c4b8583573', 34, '2026-06-19 13:48:31+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5a315aa6fc21fdde377b346cd97f6a2d24b64402a3d2d73dc10049628e1ee7c0', 48, '2026-06-19 14:30:11+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb20d1bb9150be1c1d24969f50d8e39466e88e8330480e7bc7bec637b7f442791', 34, '2026-06-19 14:30:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x7cc4268ab2a56e66d8da50128db09f3df599f00fd690a8fb91d6f33fa8073792', 34, '2026-06-19 22:58:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x3edc51f4683338fc2eb06557cb3b46c613a5c2950fda8e1b8c147a25f70a4a53', 34, '2026-06-20 16:56:47+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x27a122245ac470070431c5ed80db8cef1b7fd68b0fb91e87f91ba77afe9e7ba6', 34, '2026-06-20 20:36:49+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1cad1797b62b154691c3fb93993b70a8cd5c6610b57bbb1ab7f35873ce93d878', 34, '2026-06-20 22:07:36+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd4ec66782f525d43ec63ece11d5379e6d3f423c1683d79c3b8afc521a17995c9', 39, '2026-06-21 19:21:44+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb3af1220c4084eabf0fe507a87d90c0fc58673d5cf8c439a44c9b1168380d541', 46, '2026-06-21 23:35:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd8ace24de529304a7b10dc147574f57079a9bab5b9906690544a37887dd3a12d', 46, '2026-06-21 23:37:19+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb2f334ff6a646d4b5bf3f1c6376933cdc0557f13372529c7987b63969c5a37b4', 34, '2026-06-22 00:28:50+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb5852d9277e3281be90f29b5d9cd3cf3a3951e0476cbb3451370e030b7e6ab4f', 34, '2026-06-22 09:40:26+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x6fe0dc325f7f137129977c09aad6cf94dd7bafd871aae9dbe1e57d59a8b869a0', 39, '2026-06-22 12:48:43+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf0b6d9ec55e84a7fc19feac44e1cebbff979003ad7b6cee47786cae227893631', 37, '2026-06-24 10:20:03+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x327dbdcf0faad6df6f72e721c4c72ce1fe3166be6b9bc8042162790beabbf3ed', 46, '2026-06-24 22:20:04+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5515fa032867e2bfa50bdab7353e3891f10b95ca82993ef3a439c0e062bfb68b', 34, '2026-06-24 22:20:13+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf20d57ac8e0ab966333d538f962e3753582815363e7fab9b5de8dee584fde55a', 34, '2026-06-25 10:36:47+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xad824f273534f640547a3ec312bbe94f9057fb16026445ac306ca59f46bc1146', 34, '2026-06-25 13:11:16+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x52baa4da34cbd3911b4ac83d311c9a1a6e159c4b2a57d1ac2df5d150e3e84e44', 34, '2026-06-25 13:14:30+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1e0063577f6baf114aa47e49f5ae670290ea0becf7e1cae6f963e7865d74ff53', 37, '2026-06-25 13:16:01+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x233160fe24ac61e68614eec9b22d3623a68b0dc667975a8814ec82191614c349', 34, '2026-06-25 13:44:03+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x1491167784946529e2abe10a947e7db4c7c032cc4b043abbd5f7bc32023ea6bf', 37, '2026-06-25 13:44:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x520315ea2b34bbd7ee0c793c696ef1adb76231403136d72d12470351f082f48d', 34, '2026-06-25 13:59:57+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x540b0290bb4ee3d5d6ded1185c71f23d48334b7a7d5e18c11857cdfb289f7ad9', 47, '2026-06-25 14:00:54+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x26b761dac1631f5d7bd8f064adb95e7fcdd5bf7fb750d4a0897807dd237d9d95', 34, '2026-06-25 14:01:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2a472902ab1e10ba0d750501e39475f9a30195d154f5f3374e5ace6f5be1a894', 46, '2026-06-25 14:01:38+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x25b31e17f21c8db007e716aae9862a2aa042158f3fb19f44ec76b82a9e1dec9a', 47, '2026-06-25 14:01:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x79dc0dbb454ac9ebe1c2d0d6781f48d3c484b86a42d900c039f9f64a2c13d6c7', 34, '2026-06-25 14:10:33+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb21e6c6051f708b29532fec1f92064f478072b1b4c69a9693f28c9bc02f7f7af', 34, '2026-06-25 14:17:41+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x422c1d48bd2f92826215a0d9305206b5c3fe16f303d491a54ae1cc2bf3bc8e6e', 37, '2026-06-26 00:58:59+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x394678daed65ceacdc2a751a764ee42cbb7b5033f23c1eb22365777782131cef', 37, '2026-06-26 22:57:22+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb414042cfec46039ebe20ad250830b832110da370a2cfad9013cdb0f16708d39', 72, '2026-06-27 22:01:12+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x167b8bee45ced605f96f554ac9f8ab9c5e4c200441fe311ff485a675447e0ab6', 37, '2026-07-03 17:07:11+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x38561e59528756c0d5f2a7f1d1752199fc3a8e90796ef0a0c1655797f51ddd3f', 37, '2026-07-04 10:06:35+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4b4dccc64b60f2e0a1afd7724a0c9ad3c6c2a33fbbea0e73293110260ce3f4aa', 37, '2026-07-05 21:54:25+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfde4c434be5c33f4d389523ea87d7b0d66e24739a29addc94b6fb0f2fb13cae8', 37, '2026-07-05 22:45:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x0357762a44cb792dc96e722327a595176981637387abf01ad61dd8e5bd238f3c', 37, '2026-07-06 17:55:47+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2bea4a34c59ab229dd125bfaddb4b2eb742314f9591c2f9a4407f8bef24ad47e', 37, '2026-07-06 17:58:17+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xccc0efc5df0cf7682ee1bbff805c4a44824a6b6dc0222df5265975b0204ef214', 37, '2026-07-05 18:26:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xc2f25194028d1af039fb00a50bc69004ed5f839a8a040572f17715074680481d', 37, '2026-07-05 18:30:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x65f510db168a088f52688bbbf19f8728a7d7a8c4d4ef7afd2e46a2b38a0aaee8', 37, '2026-07-05 18:55:14+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4b763bcd43b31f518ddc7749f2db4908e5cb88dc11df990125d491b7f004c907', 37, '2026-07-05 19:06:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x44ec117c70eabf6b5a657074f6874daa46fce61b7fa154e7656f2a02c3fb6e23', 37, '2026-07-05 22:09:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9b4e60d48ceb39c5f6bc9e7fe31bf9abc57b96fc862b833b0921723d8e5e344e', 37, '2026-07-05 22:14:23+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x08970c4006738a3fd13bdedd7a6ebac66f4df4d3e2b7ecd8330de70536b076ad', 34, '2026-07-05 22:14:40+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x6f37cfe5a2a51685bf2e6e4359635344d4617a3aa27a7d7f64d3f2ef9b90b983', 37, '2026-07-05 22:18:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4dd50368d0ed399412b0df82375f76138195d97dfa8fb394fb8f8e8072b11cc1', 34, '2026-07-05 22:19:46+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb72d595ca13141ef52d60a3c4941dadc878bf90a9fbcbddf49769186765c1b07', 46, '2026-07-05 22:20:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9f0aeaab99f7ef1b2ec356b924d6dcbe0e0f039be65837f824e07769e56a93d3', 47, '2026-07-05 22:21:51+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x31a060f66f6c6ad81768057b0872778bb22e8e72198d88bdd4dfdaff5fc5401a', 34, '2026-07-05 22:22:08+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x184314282bc99335e5109b9db4eef198ceb6c1f8a327224a29fedaf26800905c', 39, '2026-07-05 22:23:50+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x9c6b253159ee9005e7fa372a96f84a68717ecbd97e97b90d694bc99e70695ff9', 34, '2026-07-05 22:24:23+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x99fe9520735f324eebbe85fc1facc04a9d6607501469d47d48c079b269982fc1', 47, '2026-07-05 22:24:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xbe7fcf832b41fbe72084cf91787d37020fe897c4a135244688940d2da9fd200f', 34, '2026-07-05 22:24:45+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5be709a5781135946312bd55c87bf1892e93a78bc4fd1a5bd7e4f706aa7ce006', 39, '2026-07-05 22:24:50+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd2a3643fab99a470495f4fc975ff47f6909fa87d4b95ebfdcd6a53fadc28a602', 34, '2026-07-05 22:24:58+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x10ad1a261fe22dd6b5efef30cb52fbedb1f5cb7864d0ac524940dd07cd10e420', 37, '2026-07-05 22:25:43+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x094015a53bf797f4fead164468a4c5d8cef2f8ba74002abbe7b071123e5481d3', 47, '2026-07-05 22:26:03+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xc639c3b67b4c7ec22a9428e4b952540755026b276c05159b997bd4e167253fdc', 37, '2026-07-06 22:28:57+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xf2903869274735c5cdf7da188b7f7d646a192e067c163e029921648948cb0fec', 46, '2026-07-06 22:38:26+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xebc93bdfbf77a6a630545c5c867ef816d0687ad96ad7a952f3a717f870056578', 47, '2026-07-06 22:38:32+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xfca3f23ee568119fa81068092664623aa60de65935dd07961800ba64130c6f8c', 39, '2026-07-06 22:40:00+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x2b03f69f6511f2be5e7d79ec1b7dc43e21e93dae02586438ea72cee4cdd028ca', 34, '2026-07-06 22:42:07+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xd531c4d8c3095b4840d5355174a0d7ae22234347449dafd7cc783af3e88340e5', 37, '2026-07-08 13:10:52+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x14b5054c23c5beae525cfd1504fac2803e019bc1760b990c41120d582ac3e86c', 46, '2026-07-08 13:21:16+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x36004fac36b146db987caf9212f172ca889f4af423085a94c40d6c24a130c100', 39, '2026-07-08 13:21:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x62ae0c54a06e0d4d6968d0cfbbc071b2c8b17568509ff06e16d49ba1d4a0e91b', 46, '2026-07-08 21:06:09+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xeb3836a83697ff0fb2aa7ccb697fe29ae5fce73171e5053ae2f916fc01bd6641', 37, '2026-07-08 21:06:51+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x30e6b7479c373954de67dc56bb117293cfec8928132ceca15aa2fd2e4ffee53b', 48, '2026-07-08 21:07:23+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x0fe59d77877f8cb98363240feb7d6015ef72c98d9a1dbf5efdc59235e323e0a0', 39, '2026-07-08 21:07:34+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x7d7da759269e5d8c214c717c361dab7dfe3e258ce3a5b63355f56a7489638c3f', 39, '2026-07-09 16:43:50+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xb89618ec197c5c158265b4f667d0be90da1daca014344d7608e29efff2812295', 46, '2026-07-10 11:27:20+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xe941a0d668af76a5545019db896eeec933594bc8e70282027d0c0c64cbb39b98', 46, '2026-07-10 11:51:34+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5b68c4e4149631be1940b79c5f56bc76e9fb951490eab9310b5b7b4311b78b5f', 39, '2026-07-10 12:01:40+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\xe60fa4418ce75598c20a51730065b13ade88f2243a31098d0df15de753aa4201', 46, '2026-07-10 12:02:42+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x03a7d0c16a40be38a34ad36d565312c01b3f74a95a9279aa5bacdfe7ff28dcc9', 39, '2026-07-13 17:10:38+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4af0e9a0d0da1035a2ca7ef3d91d366dfc1f2d268588b0fc38787e2211bfdb08', 39, '2026-07-13 17:41:17+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x5e438de7e2b8898309074a8422c912fee0b6813c1438c415d0081184f5eece47', 37, '2026-07-13 21:11:36+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x4a82cd93d768b3db919c190f186d1d221a2942776f76047a40d8ad97d9587c11', 39, '2026-07-13 21:12:13+07', 'authentication');
INSERT INTO public.tokens (hash, user_id, expiry, scope) VALUES ('\x98fd81c3a98cb0e5f647a12b5f1289d520e8d09d6fb6713d98bed282df890dfe', 46, '2026-07-13 21:13:14+07', 'authentication');


ALTER TABLE public.tokens ENABLE TRIGGER ALL;

--
-- Data for Name: tracking; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.tracking DISABLE TRIGGER ALL;

INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5329, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-07 13:58:00.885145+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5330, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-07 13:58:11.125496+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5331, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-07 13:58:21.260664+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5332, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-07 13:58:31.397485+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5333, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-07 13:58:41.535306+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5334, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-07 13:58:51.673761+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5335, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-07 13:59:01.913106+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5336, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-07 13:59:12.146577+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5337, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-07 13:59:22.286015+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5338, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 13:59:32.429147+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5339, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 13:59:42.566784+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5340, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 13:59:52.704354+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5341, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 14:00:02.84146+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5342, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 14:00:12.980632+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5343, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 14:00:23.113846+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5344, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 14:00:33.279072+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5345, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 14:00:43.416219+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5346, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-07 14:00:53.554036+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5347, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-07 14:01:04.638719+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5348, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-07 14:01:15.138422+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5349, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 17:47:03.613004+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5350, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 17:47:14.019566+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5351, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 17:47:24.085899+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5352, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 17:47:34.175175+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5353, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 17:47:44.391549+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5354, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 17:47:54.617681+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5355, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 17:48:04.8347+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5356, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 17:48:15.06418+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5357, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-07 17:48:25.293672+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5358, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-07 17:48:35.386866+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5359, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-07 17:48:45.447203+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5360, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-07 17:48:55.508042+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5361, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-07 17:49:05.603139+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5362, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-07 17:49:15.680654+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5363, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-07 17:49:25.761254+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5364, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-07 17:49:35.91887+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5365, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-07 17:49:46.026318+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5366, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-07 17:49:56.159935+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5367, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-07 17:50:06.217722+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5368, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-07 17:50:16.278236+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5369, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-07 17:50:26.440549+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5370, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 17:50:36.582784+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5371, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 17:50:46.793419+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5372, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 17:50:57.758741+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5373, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 17:51:08.559522+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5374, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 17:51:53.498955+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5375, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 17:52:37.787309+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5376, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 17:53:21.84329+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5377, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 18:02:45.295383+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5378, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 18:03:19.790185+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5379, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 18:03:29.88453+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5380, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 18:03:39.955128+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5381, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 18:03:50.045979+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5382, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 18:04:00.127752+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5383, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 18:04:10.240515+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5384, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 18:04:20.341399+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5385, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 18:04:30.467824+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5386, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-07 18:04:40.54824+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5387, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-07 18:04:50.622253+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5388, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-07 18:05:00.707378+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5389, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-07 18:05:10.781144+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5390, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-07 18:05:20.934138+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5391, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-07 18:05:31.060772+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5392, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-07 18:05:41.137683+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5393, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-07 18:05:51.200446+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5394, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-07 18:06:01.277961+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5395, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-07 18:06:11.771141+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5396, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-07 18:06:21.868153+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5397, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-07 18:06:31.951574+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5398, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-07 18:06:42.03749+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5399, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 18:06:52.23343+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5400, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 18:07:02.477301+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5401, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 18:07:12.710035+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5402, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 18:07:22.869649+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5403, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 18:07:33.063201+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5404, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 18:07:43.643128+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5405, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 18:07:54.004619+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5406, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 18:08:04.093151+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5407, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-07 18:08:14.174662+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5408, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-07 18:08:24.250761+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5409, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-07 18:08:34.35663+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5410, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-07 18:08:44.512285+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5411, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-07 18:08:54.732705+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5412, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-07 18:09:04.940875+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5413, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-07 18:09:15.019882+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5414, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-07 18:09:25.097417+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5415, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-07 18:09:35.349691+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5416, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-07 18:09:45.473209+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5417, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-07 18:09:55.551117+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5418, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-07 18:10:05.629451+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5419, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-07 18:10:15.755533+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5420, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 18:10:25.830416+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5421, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 18:10:35.89565+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5422, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-07 18:10:45.971033+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5423, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-07 18:10:56.055184+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5424, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-07 18:11:06.124266+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5425, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-07 18:11:16.197429+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5426, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-07 18:11:26.263777+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5427, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-07 18:11:36.353125+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5428, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-07 18:11:46.473181+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5429, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-07 18:11:56.591875+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5430, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-07 18:12:06.718099+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5431, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-07 18:12:16.7878+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5432, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-07 18:12:26.863728+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5433, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-07 18:12:36.940143+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5434, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-07 18:12:47.074311+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5435, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-07 18:12:57.18391+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5436, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-07 18:13:07.3101+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5437, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-07 18:13:18.106511+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5438, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-07 18:13:38.934656+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5439, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-07 18:13:55.557136+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5440, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-07 18:14:06.560029+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5441, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-07 18:14:17.563368+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5442, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-07 18:19:52.001079+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5443, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 11:27:45.600331+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5444, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 11:27:55.738579+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5445, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-09 11:28:05.799219+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5446, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-09 11:28:15.845753+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5447, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-09 11:28:25.929023+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5448, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-09 11:28:36.070051+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5449, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-09 11:28:46.147566+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5450, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-09 11:28:56.330523+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5451, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-09 11:29:06.397755+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5452, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-09 11:29:16.54294+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5453, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-09 11:29:26.600029+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5454, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-09 11:29:36.659797+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5455, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-09 11:29:46.721439+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5456, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-09 11:29:57.017679+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5457, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-09 11:30:07.240367+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5458, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-09 11:30:17.458046+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5459, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-09 11:30:27.579362+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5460, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-09 11:30:37.795669+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5461, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-09 11:30:47.874024+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5462, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-09 11:30:57.969557+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5463, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-09 11:31:08.019053+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5464, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 11:31:18.087498+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5465, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 11:31:28.174026+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5466, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-09 11:31:38.230676+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5467, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-09 11:31:48.340132+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5468, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-09 11:31:58.415983+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5469, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-09 11:32:08.479322+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5470, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-09 11:32:18.635051+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5471, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-09 11:32:28.764605+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5472, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-09 11:32:38.830501+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5473, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-09 11:32:48.88805+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5474, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-09 11:32:59.023845+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5475, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-09 11:33:09.110753+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5476, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-09 11:33:19.165096+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5477, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-09 11:33:29.292588+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5478, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-09 11:33:39.443839+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5479, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-09 11:33:49.555471+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5480, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-09 11:33:59.642704+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5481, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-09 11:34:09.708339+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5482, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-09 11:34:19.780616+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5483, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-09 11:34:29.855848+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5484, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-09 11:34:39.941826+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5485, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 11:34:50.005109+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5486, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 11:35:00.072711+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5487, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-09 11:35:10.19128+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5488, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-09 11:35:20.357945+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5489, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-09 11:35:30.404785+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5490, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-09 11:35:40.52546+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5491, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-09 11:35:50.60072+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5492, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-09 11:36:00.678841+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5493, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-09 11:36:10.75734+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5494, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-09 11:36:20.823628+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5495, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-09 11:36:30.893681+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5496, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-09 11:36:40.953504+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5497, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-09 11:36:51.092341+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5498, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-09 11:37:01.165379+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5499, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-09 11:37:11.246477+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5500, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-09 11:37:21.362242+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5501, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-09 11:37:31.473771+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5502, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-09 11:37:42.942099+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5503, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-09 11:37:53.6123+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5504, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-09 11:38:21.806036+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5505, 175, 0.12770046445395156, 110.60128011718128, 20, 5, '2026-07-09 11:38:32.822078+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5506, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 11:39:07.787409+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5507, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 11:47:02.25044+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5508, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 12:03:25.298634+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5509, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 12:03:35.542864+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5510, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-09 12:03:45.733446+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5511, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-09 12:03:55.815788+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5512, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-09 12:04:05.884731+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5513, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-09 12:04:15.965303+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5514, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-09 12:04:26.036411+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5515, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-09 12:04:36.107287+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5516, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-09 12:04:46.161169+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5517, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-09 12:04:56.257663+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5518, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-09 12:05:06.326828+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5519, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-09 12:05:16.404243+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5520, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-09 12:05:26.535751+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5521, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-09 12:05:36.6951+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5522, 175, 0.1261925391225784, 110.60166676565105, 20, 5, '2026-07-09 12:05:47.383881+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5523, 175, 0.12651474539733873, 110.60075169760592, 20, 5, '2026-07-09 12:05:58.3539+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5524, 175, 0.1266178514044249, 110.60071303275892, 20, 5, '2026-07-09 12:06:09.619783+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5525, 175, 0.12686272816959912, 110.60095791012311, 20, 5, '2026-07-09 12:06:20.398042+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5526, 175, 0.12719782268660207, 110.6010996812287, 20, 5, '2026-07-09 12:06:31.361921+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5527, 175, 0.1272880404404282, 110.60116412264034, 20, 5, '2026-07-09 12:07:51.984212+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5528, 175, 0.12338290023859251, 110.60558480347815, 20, 5, '2026-07-09 12:09:30.008478+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5529, 175, 0.12311224693583053, 110.6054043675256, 20, 5, '2026-07-09 12:09:40.12422+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5530, 175, 0.12300914091500435, 110.60478572997395, 20, 5, '2026-07-09 12:09:50.192278+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5531, 175, 0.12285448188302084, 110.60408976272834, 20, 5, '2026-07-09 12:10:00.269392+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5532, 175, 0.12275137586120231, 110.60349690174134, 20, 5, '2026-07-09 12:10:10.359458+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5533, 175, 0.12269982285014676, 110.60309736498924, 20, 5, '2026-07-09 12:10:20.485638+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5534, 175, 0.12295758790444496, 110.60299425873063, 20, 5, '2026-07-09 12:10:30.617234+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5535, 175, 0.123318458976287, 110.6028782641897, 20, 5, '2026-07-09 12:10:41.355718+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5536, 175, 0.12378243606144913, 110.60272360480178, 20, 5, '2026-07-09 12:10:52.359644+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5537, 175, 0.12438818391041155, 110.60260761026085, 20, 5, '2026-07-09 12:11:03.457878+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5538, 175, 0.12487793748027519, 110.60251739228457, 20, 5, '2026-07-09 12:11:14.711067+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5539, 175, 0.1252774732805117, 110.60241428602595, 20, 5, '2026-07-09 12:11:25.446477+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5540, 175, 0.12553523830968352, 110.60227251492037, 20, 5, '2026-07-09 12:11:36.400864+07');
INSERT INTO public.tracking (id, pengiriman_id, latitude, longitude, speed, accuracy, created_at) VALUES (5541, 175, 0.12594766235108124, 110.6020405258385, 20, 5, '2026-07-09 12:14:24.403435+07');


ALTER TABLE public.tracking ENABLE TRIGGER ALL;

--
-- Data for Name: users_permissions; Type: TABLE DATA; Schema: public; Owner: rian
--

ALTER TABLE public.users_permissions DISABLE TRIGGER ALL;

INSERT INTO public.users_permissions (user_id, permission_id) VALUES (37, 1);


ALTER TABLE public.users_permissions ENABLE TRIGGER ALL;

--
-- Name: alokasi_harian_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.alokasi_harian_id_seq', 9, true);


--
-- Name: driver_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.driver_id_seq', 9, true);


--
-- Name: kecamatan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.kecamatan_id_seq', 15, true);


--
-- Name: kelurahan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.kelurahan_id_seq', 164, true);


--
-- Name: pedagang_lokal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.pedagang_lokal_id_seq', 20, true);


--
-- Name: pending_sppg_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.pending_sppg_id_seq', 1, false);


--
-- Name: pengeluaran_harian_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.pengeluaran_harian_id_seq', 20, true);


--
-- Name: pengiriman_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.pengiriman_id_seq', 185, true);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.permissions_id_seq', 2, true);


--
-- Name: posyandu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.posyandu_id_seq', 10, true);


--
-- Name: produksi_harian_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.produksi_harian_id_seq', 3, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.roles_id_seq', 4, true);


--
-- Name: sekolah_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.sekolah_id_seq', 59, true);


--
-- Name: sppg_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.sppg_id_seq', 18, true);


--
-- Name: sppg_invitations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.sppg_invitations_id_seq', 9, true);


--
-- Name: tracking_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.tracking_id_seq', 5541, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rian
--

SELECT pg_catalog.setval('public.users_id_seq', 87, true);


--
-- PostgreSQL database dump complete
--

\unrestrict YIGuEs4fMvYa0rwCQjyqwIH3LPqeu8SGIzxVS9aSRmfHBDnU0Qm5SRqgSP5TiIA

