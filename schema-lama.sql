--
-- PostgreSQL database dump
--

\restrict eXcn5F6zswsO6Wu3evG9gIrBUO7Ffg2LiIc9UdLQUQ5qfXEayZtO7c5G2LkHRlr

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
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: alokasi_harian; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.alokasi_harian (
    id bigint NOT NULL,
    sppg_id bigint NOT NULL,
    tanggal date NOT NULL,
    jumlah bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.alokasi_harian OWNER TO rian;

--
-- Name: alokasi_harian_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.alokasi_harian_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.alokasi_harian_id_seq OWNER TO rian;

--
-- Name: alokasi_harian_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.alokasi_harian_id_seq OWNED BY public.alokasi_harian.id;


--
-- Name: drivers; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.drivers (
    id bigint CONSTRAINT driver_id_not_null NOT NULL,
    created_at timestamp with time zone DEFAULT now() CONSTRAINT driver_created_at_not_null NOT NULL,
    user_id bigint CONSTRAINT driver_user_id_not_null NOT NULL,
    sppg_id bigint CONSTRAINT driver_sppg_id_not_null NOT NULL,
    nama text CONSTRAINT driver_nama_not_null NOT NULL,
    nomor_telepon text CONSTRAINT driver_nomor_telepon_not_null NOT NULL,
    status_aktif boolean DEFAULT true CONSTRAINT driver_is_active_not_null NOT NULL,
    version integer DEFAULT 1 CONSTRAINT driver_version_not_null NOT NULL,
    CONSTRAINT driver_nama_check CHECK ((char_length(TRIM(BOTH FROM nama)) > 0)),
    CONSTRAINT driver_nomor_telepon_check CHECK ((char_length(TRIM(BOTH FROM nomor_telepon)) > 0)),
    CONSTRAINT driver_version_check CHECK ((version >= 1))
);


ALTER TABLE public.drivers OWNER TO rian;

--
-- Name: driver_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.driver_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.driver_id_seq OWNER TO rian;

--
-- Name: driver_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.driver_id_seq OWNED BY public.drivers.id;


--
-- Name: kecamatan; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.kecamatan (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.kecamatan OWNER TO rian;

--
-- Name: kecamatan_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.kecamatan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kecamatan_id_seq OWNER TO rian;

--
-- Name: kecamatan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.kecamatan_id_seq OWNED BY public.kecamatan.id;


--
-- Name: kelurahan; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.kelurahan (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    kecamatan_id bigint NOT NULL,
    name text NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.kelurahan OWNER TO rian;

--
-- Name: kelurahan_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.kelurahan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kelurahan_id_seq OWNER TO rian;

--
-- Name: kelurahan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.kelurahan_id_seq OWNED BY public.kelurahan.id;


--
-- Name: pedagang_lokal; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.pedagang_lokal (
    id bigint NOT NULL,
    nama text NOT NULL,
    alamat text NOT NULL,
    no_hp text NOT NULL,
    longitude double precision NOT NULL,
    latitude double precision NOT NULL,
    jenis_produk text NOT NULL,
    sppg_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.pedagang_lokal OWNER TO rian;

--
-- Name: pedagang_lokal_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.pedagang_lokal_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pedagang_lokal_id_seq OWNER TO rian;

--
-- Name: pedagang_lokal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.pedagang_lokal_id_seq OWNED BY public.pedagang_lokal.id;


--
-- Name: pending_sppg; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.pending_sppg (
    id bigint NOT NULL,
    name text NOT NULL,
    email public.citext NOT NULL,
    password_hash bytea NOT NULL,
    nama text NOT NULL,
    alamat text NOT NULL,
    sosmed_url text[],
    kepala_sppg text,
    nomor_telepon text,
    latitude double precision,
    longitude double precision,
    kapasitas_porsi integer DEFAULT 0 NOT NULL,
    kecamatan_id bigint,
    kelurahan_id bigint,
    sppg_invitations_id bigint NOT NULL,
    status text DEFAULT 'pending'::text NOT NULL,
    approved_by bigint,
    approved_at timestamp(0) with time zone,
    rejected_reason text,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    CONSTRAINT pending_sppg_alamat_length_check CHECK ((char_length(alamat) <= 500)),
    CONSTRAINT pending_sppg_kapasitas_porsi_check CHECK ((kapasitas_porsi >= 0)),
    CONSTRAINT pending_sppg_kepala_sppg_length_check CHECK (((kepala_sppg IS NULL) OR (char_length(kepala_sppg) <= 200))),
    CONSTRAINT pending_sppg_latitude_check CHECK (((latitude IS NULL) OR ((latitude >= ('-90'::integer)::double precision) AND (latitude <= (90)::double precision)))),
    CONSTRAINT pending_sppg_longitude_check CHECK (((longitude IS NULL) OR ((longitude >= ('-180'::integer)::double precision) AND (longitude <= (180)::double precision)))),
    CONSTRAINT pending_sppg_nama_length_check CHECK ((char_length(nama) <= 200)),
    CONSTRAINT pending_sppg_nomor_telepon_length_check CHECK (((nomor_telepon IS NULL) OR (char_length(nomor_telepon) <= 20))),
    CONSTRAINT pending_sppg_status_check CHECK ((status = ANY (ARRAY['pending'::text, 'approved'::text, 'rejected'::text])))
);


ALTER TABLE public.pending_sppg OWNER TO rian;

--
-- Name: pending_sppg_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.pending_sppg_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pending_sppg_id_seq OWNER TO rian;

--
-- Name: pending_sppg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.pending_sppg_id_seq OWNED BY public.pending_sppg.id;


--
-- Name: pengeluaran_harian; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.pengeluaran_harian (
    id bigint NOT NULL,
    alokasi_harian_id bigint NOT NULL,
    produk character varying(255) NOT NULL,
    jumlah bigint NOT NULL,
    satuan character varying(50) NOT NULL,
    harga_satuan bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    subtotal bigint GENERATED ALWAYS AS ((jumlah * harga_satuan)) STORED,
    pedagang_lokal_id bigint
);


ALTER TABLE public.pengeluaran_harian OWNER TO rian;

--
-- Name: pengeluaran_harian_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.pengeluaran_harian_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pengeluaran_harian_id_seq OWNER TO rian;

--
-- Name: pengeluaran_harian_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.pengeluaran_harian_id_seq OWNED BY public.pengeluaran_harian.id;


--
-- Name: pengiriman; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.pengiriman (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    sppg_id bigint NOT NULL,
    driver_id bigint,
    tujuan_type text NOT NULL,
    tujuan_id bigint NOT NULL,
    status text DEFAULT 'menunggu'::text NOT NULL,
    waktu_berangkat timestamp with time zone,
    waktu_selesai timestamp with time zone,
    version integer DEFAULT 1 NOT NULL,
    CONSTRAINT pengiriman_status_check CHECK ((status = ANY (ARRAY['menunggu'::text, 'berangkat'::text, 'sampai'::text, 'dibatalkan'::text]))),
    CONSTRAINT pengiriman_tujuan_type_check CHECK ((tujuan_type = ANY (ARRAY['sekolah'::text, 'posyandu'::text]))),
    CONSTRAINT pengiriman_version_check CHECK ((version >= 1))
);


ALTER TABLE public.pengiriman OWNER TO rian;

--
-- Name: pengiriman_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.pengiriman_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pengiriman_id_seq OWNER TO rian;

--
-- Name: pengiriman_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.pengiriman_id_seq OWNED BY public.pengiriman.id;


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    code text NOT NULL
);


ALTER TABLE public.permissions OWNER TO rian;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.permissions_id_seq OWNER TO rian;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: posyandu; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.posyandu (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    sppg_id bigint NOT NULL,
    nama text NOT NULL,
    alamat text NOT NULL,
    latitude double precision,
    longitude double precision,
    jumlah_balita integer DEFAULT 0 NOT NULL,
    jumlah_ibu_hamil integer DEFAULT 0 NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    kecamatan_id bigint NOT NULL,
    kelurahan_id bigint NOT NULL,
    CONSTRAINT posyandu_alamat_length_check CHECK ((char_length(alamat) <= 500)),
    CONSTRAINT posyandu_jumlah_balita_check CHECK ((jumlah_balita >= 0)),
    CONSTRAINT posyandu_jumlah_ibu_hamil_check CHECK ((jumlah_ibu_hamil >= 0)),
    CONSTRAINT posyandu_latitude_check CHECK (((latitude IS NULL) OR ((latitude >= ('-90'::integer)::double precision) AND (latitude <= (90)::double precision)))),
    CONSTRAINT posyandu_longitude_check CHECK (((longitude IS NULL) OR ((longitude >= ('-180'::integer)::double precision) AND (longitude <= (180)::double precision)))),
    CONSTRAINT posyandu_nama_length_check CHECK ((char_length(nama) <= 200))
);


ALTER TABLE public.posyandu OWNER TO rian;

--
-- Name: posyandu_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.posyandu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.posyandu_id_seq OWNER TO rian;

--
-- Name: posyandu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.posyandu_id_seq OWNED BY public.posyandu.id;


--
-- Name: produksi_harian; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.produksi_harian (
    id bigint NOT NULL,
    sppg_id bigint NOT NULL,
    tanggal date NOT NULL,
    waktu_mulai timestamp without time zone NOT NULL,
    estimasi_waktu_selesai timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.produksi_harian OWNER TO rian;

--
-- Name: produksi_harian_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.produksi_harian_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.produksi_harian_id_seq OWNER TO rian;

--
-- Name: produksi_harian_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.produksi_harian_id_seq OWNED BY public.produksi_harian.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text NOT NULL,
    permissions text[] DEFAULT '{}'::text[] NOT NULL
);


ALTER TABLE public.roles OWNER TO rian;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO rian;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO rian;

--
-- Name: sekolah; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.sekolah (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    nama text NOT NULL,
    tingkat text NOT NULL,
    jumlah_siswa integer NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    alamat text DEFAULT ''::text NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    sppg_id bigint,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    kecamatan_id bigint NOT NULL,
    kelurahan_id bigint NOT NULL,
    CONSTRAINT sekolah_alamat_length_check CHECK ((char_length(alamat) <= 500)),
    CONSTRAINT sekolah_jumlah_siswa_check CHECK ((jumlah_siswa > 0)),
    CONSTRAINT sekolah_latitude_check CHECK (((latitude >= ('-90'::integer)::double precision) AND (latitude <= (90)::double precision))),
    CONSTRAINT sekolah_longitude_check CHECK (((longitude >= ('-180'::integer)::double precision) AND (longitude <= (180)::double precision))),
    CONSTRAINT sekolah_nama_length_check CHECK ((char_length(nama) <= 200)),
    CONSTRAINT sekolah_tingkat_check CHECK ((tingkat = ANY (ARRAY['SD'::text, 'SMP'::text, 'SMA'::text])))
);


ALTER TABLE public.sekolah OWNER TO rian;

--
-- Name: sekolah_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.sekolah_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sekolah_id_seq OWNER TO rian;

--
-- Name: sekolah_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.sekolah_id_seq OWNED BY public.sekolah.id;


--
-- Name: sppg; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.sppg (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    nama text,
    alamat text,
    sosmed_url text[],
    kepala_sppg text,
    nomor_telepon text,
    email text,
    latitude double precision,
    longitude double precision,
    kapasitas_porsi integer DEFAULT 0 NOT NULL,
    status_aktif boolean DEFAULT true NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    kecamatan_id bigint,
    kelurahan_id bigint,
    CONSTRAINT sppg_alamat_length_check CHECK ((char_length(alamat) <= 500)),
    CONSTRAINT sppg_email_length_check CHECK (((email IS NULL) OR (char_length(email) <= 255))),
    CONSTRAINT sppg_kapasitas_porsi_check CHECK ((kapasitas_porsi >= 0)),
    CONSTRAINT sppg_kepala_sppg_length_check CHECK ((char_length(kepala_sppg) <= 200)),
    CONSTRAINT sppg_latitude_check CHECK (((latitude IS NULL) OR ((latitude >= ('-90'::integer)::double precision) AND (latitude <= (90)::double precision)))),
    CONSTRAINT sppg_longitude_check CHECK (((longitude IS NULL) OR ((longitude >= ('-180'::integer)::double precision) AND (longitude <= (180)::double precision)))),
    CONSTRAINT sppg_nama_length_check CHECK ((char_length(nama) <= 200)),
    CONSTRAINT sppg_nomor_telepon_length_check CHECK (((nomor_telepon IS NULL) OR (char_length(nomor_telepon) <= 20)))
);


ALTER TABLE public.sppg OWNER TO rian;

--
-- Name: sppg_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.sppg_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sppg_id_seq OWNER TO rian;

--
-- Name: sppg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.sppg_id_seq OWNED BY public.sppg.id;


--
-- Name: sppg_invitations; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.sppg_invitations (
    id bigint NOT NULL,
    token text NOT NULL,
    nama_sppg text NOT NULL,
    expires_at timestamp with time zone,
    used_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.sppg_invitations OWNER TO rian;

--
-- Name: sppg_invitations_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.sppg_invitations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sppg_invitations_id_seq OWNER TO rian;

--
-- Name: sppg_invitations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.sppg_invitations_id_seq OWNED BY public.sppg_invitations.id;


--
-- Name: tokens; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.tokens (
    hash bytea NOT NULL,
    user_id bigint NOT NULL,
    expiry timestamp(0) with time zone NOT NULL,
    scope text NOT NULL
);


ALTER TABLE public.tokens OWNER TO rian;

--
-- Name: tracking; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.tracking (
    id bigint NOT NULL,
    pengiriman_id bigint NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    speed double precision,
    accuracy double precision,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tracking OWNER TO rian;

--
-- Name: tracking_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.tracking_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tracking_id_seq OWNER TO rian;

--
-- Name: tracking_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.tracking_id_seq OWNED BY public.tracking.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    email public.citext NOT NULL,
    password_hash bytea NOT NULL,
    activated boolean NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    role_id bigint NOT NULL,
    last_active_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO rian;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: rian
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO rian;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rian
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users_permissions; Type: TABLE; Schema: public; Owner: rian
--

CREATE TABLE public.users_permissions (
    user_id bigint NOT NULL,
    permission_id bigint NOT NULL
);


ALTER TABLE public.users_permissions OWNER TO rian;

--
-- Name: alokasi_harian id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.alokasi_harian ALTER COLUMN id SET DEFAULT nextval('public.alokasi_harian_id_seq'::regclass);


--
-- Name: drivers id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers ALTER COLUMN id SET DEFAULT nextval('public.driver_id_seq'::regclass);


--
-- Name: kecamatan id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kecamatan ALTER COLUMN id SET DEFAULT nextval('public.kecamatan_id_seq'::regclass);


--
-- Name: kelurahan id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kelurahan ALTER COLUMN id SET DEFAULT nextval('public.kelurahan_id_seq'::regclass);


--
-- Name: pedagang_lokal id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pedagang_lokal ALTER COLUMN id SET DEFAULT nextval('public.pedagang_lokal_id_seq'::regclass);


--
-- Name: pending_sppg id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg ALTER COLUMN id SET DEFAULT nextval('public.pending_sppg_id_seq'::regclass);


--
-- Name: pengeluaran_harian id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengeluaran_harian ALTER COLUMN id SET DEFAULT nextval('public.pengeluaran_harian_id_seq'::regclass);


--
-- Name: pengiriman id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengiriman ALTER COLUMN id SET DEFAULT nextval('public.pengiriman_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: posyandu id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.posyandu ALTER COLUMN id SET DEFAULT nextval('public.posyandu_id_seq'::regclass);


--
-- Name: produksi_harian id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.produksi_harian ALTER COLUMN id SET DEFAULT nextval('public.produksi_harian_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: sekolah id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sekolah ALTER COLUMN id SET DEFAULT nextval('public.sekolah_id_seq'::regclass);


--
-- Name: sppg id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg ALTER COLUMN id SET DEFAULT nextval('public.sppg_id_seq'::regclass);


--
-- Name: sppg_invitations id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg_invitations ALTER COLUMN id SET DEFAULT nextval('public.sppg_invitations_id_seq'::regclass);


--
-- Name: tracking id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.tracking ALTER COLUMN id SET DEFAULT nextval('public.tracking_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: alokasi_harian alokasi_harian_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.alokasi_harian
    ADD CONSTRAINT alokasi_harian_pkey PRIMARY KEY (id);


--
-- Name: alokasi_harian alokasi_harian_sppg_id_tanggal_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.alokasi_harian
    ADD CONSTRAINT alokasi_harian_sppg_id_tanggal_key UNIQUE (sppg_id, tanggal);


--
-- Name: drivers driver_nomor_telepon_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT driver_nomor_telepon_key UNIQUE (nomor_telepon);


--
-- Name: drivers driver_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT driver_pkey PRIMARY KEY (id);


--
-- Name: drivers driver_user_id_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT driver_user_id_key UNIQUE (user_id);


--
-- Name: kecamatan kecamatan_name_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kecamatan
    ADD CONSTRAINT kecamatan_name_key UNIQUE (name);


--
-- Name: kecamatan kecamatan_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kecamatan
    ADD CONSTRAINT kecamatan_pkey PRIMARY KEY (id);


--
-- Name: kelurahan kelurahan_kecamatan_id_name_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kelurahan
    ADD CONSTRAINT kelurahan_kecamatan_id_name_key UNIQUE (kecamatan_id, name);


--
-- Name: kelurahan kelurahan_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kelurahan
    ADD CONSTRAINT kelurahan_pkey PRIMARY KEY (id);


--
-- Name: pedagang_lokal pedagang_lokal_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pedagang_lokal
    ADD CONSTRAINT pedagang_lokal_pkey PRIMARY KEY (id);


--
-- Name: pending_sppg pending_sppg_email_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT pending_sppg_email_key UNIQUE (email);


--
-- Name: pending_sppg pending_sppg_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT pending_sppg_pkey PRIMARY KEY (id);


--
-- Name: pengeluaran_harian pengeluaran_harian_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengeluaran_harian
    ADD CONSTRAINT pengeluaran_harian_pkey PRIMARY KEY (id);


--
-- Name: pengiriman pengiriman_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengiriman
    ADD CONSTRAINT pengiriman_pkey PRIMARY KEY (id);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: posyandu posyandu_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.posyandu
    ADD CONSTRAINT posyandu_pkey PRIMARY KEY (id);


--
-- Name: produksi_harian produksi_harian_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.produksi_harian
    ADD CONSTRAINT produksi_harian_pkey PRIMARY KEY (id);


--
-- Name: produksi_harian produksi_harian_sppg_id_tanggal_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.produksi_harian
    ADD CONSTRAINT produksi_harian_sppg_id_tanggal_key UNIQUE (sppg_id, tanggal);


--
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sekolah sekolah_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sekolah
    ADD CONSTRAINT sekolah_pkey PRIMARY KEY (id);


--
-- Name: sppg_invitations sppg_invitations_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg_invitations
    ADD CONSTRAINT sppg_invitations_pkey PRIMARY KEY (id);


--
-- Name: sppg_invitations sppg_invitations_token_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg_invitations
    ADD CONSTRAINT sppg_invitations_token_key UNIQUE (token);


--
-- Name: sppg sppg_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg
    ADD CONSTRAINT sppg_pkey PRIMARY KEY (id);


--
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (hash);


--
-- Name: tracking tracking_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.tracking
    ADD CONSTRAINT tracking_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users_permissions users_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_pkey PRIMARY KEY (user_id, permission_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_drivers_sppg_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_drivers_sppg_id ON public.drivers USING btree (sppg_id);


--
-- Name: idx_drivers_user_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_drivers_user_id ON public.drivers USING btree (user_id);


--
-- Name: idx_pedagang_lokal_jenis_produk; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pedagang_lokal_jenis_produk ON public.pedagang_lokal USING btree (jenis_produk);


--
-- Name: idx_pedagang_lokal_sppg; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pedagang_lokal_sppg ON public.pedagang_lokal USING btree (sppg_id);


--
-- Name: idx_pending_sppg_kecamatan_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pending_sppg_kecamatan_id ON public.pending_sppg USING btree (kecamatan_id);


--
-- Name: idx_pending_sppg_kelurahan_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pending_sppg_kelurahan_id ON public.pending_sppg USING btree (kelurahan_id);


--
-- Name: idx_pengiriman_driver_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pengiriman_driver_id ON public.pengiriman USING btree (driver_id);


--
-- Name: idx_pengiriman_sppg_id; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pengiriman_sppg_id ON public.pengiriman USING btree (sppg_id);


--
-- Name: idx_pengiriman_status; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pengiriman_status ON public.pengiriman USING btree (status);


--
-- Name: idx_pengiriman_tujuan; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX idx_pengiriman_tujuan ON public.pengiriman USING btree (tujuan_type, tujuan_id);


--
-- Name: kelurahan_kecamatan_id_idx; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX kelurahan_kecamatan_id_idx ON public.kelurahan USING btree (kecamatan_id);


--
-- Name: posyandu_sppg_id_idx; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX posyandu_sppg_id_idx ON public.posyandu USING btree (sppg_id);


--
-- Name: sekolah_nama_trgm_idx; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX sekolah_nama_trgm_idx ON public.sekolah USING gin (nama public.gin_trgm_ops);


--
-- Name: sekolah_sppg_id_idx; Type: INDEX; Schema: public; Owner: rian
--

CREATE INDEX sekolah_sppg_id_idx ON public.sekolah USING btree (sppg_id);


--
-- Name: alokasi_harian alokasi_harian_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.alokasi_harian
    ADD CONSTRAINT alokasi_harian_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id);


--
-- Name: drivers driver_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT driver_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id) ON DELETE RESTRICT;


--
-- Name: drivers driver_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT driver_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: pending_sppg fk_pending_sppg_approved_by; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT fk_pending_sppg_approved_by FOREIGN KEY (approved_by) REFERENCES public.users(id);


--
-- Name: pending_sppg fk_pending_sppg_invitation; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT fk_pending_sppg_invitation FOREIGN KEY (sppg_invitations_id) REFERENCES public.sppg_invitations(id);


--
-- Name: pending_sppg fk_pending_sppg_kecamatan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT fk_pending_sppg_kecamatan FOREIGN KEY (kecamatan_id) REFERENCES public.kecamatan(id);


--
-- Name: pending_sppg fk_pending_sppg_kelurahan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pending_sppg
    ADD CONSTRAINT fk_pending_sppg_kelurahan FOREIGN KEY (kelurahan_id) REFERENCES public.kelurahan(id);


--
-- Name: posyandu fk_posyandu_kecamatan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.posyandu
    ADD CONSTRAINT fk_posyandu_kecamatan FOREIGN KEY (kecamatan_id) REFERENCES public.kecamatan(id);


--
-- Name: posyandu fk_posyandu_kelurahan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.posyandu
    ADD CONSTRAINT fk_posyandu_kelurahan FOREIGN KEY (kelurahan_id) REFERENCES public.kelurahan(id);


--
-- Name: sekolah fk_sekolah_kecamatan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sekolah
    ADD CONSTRAINT fk_sekolah_kecamatan FOREIGN KEY (kecamatan_id) REFERENCES public.kecamatan(id);


--
-- Name: sekolah fk_sekolah_kelurahan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sekolah
    ADD CONSTRAINT fk_sekolah_kelurahan FOREIGN KEY (kelurahan_id) REFERENCES public.kelurahan(id);


--
-- Name: sppg fk_sppg_kecamatan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg
    ADD CONSTRAINT fk_sppg_kecamatan FOREIGN KEY (kecamatan_id) REFERENCES public.kecamatan(id);


--
-- Name: sppg fk_sppg_kelurahan; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg
    ADD CONSTRAINT fk_sppg_kelurahan FOREIGN KEY (kelurahan_id) REFERENCES public.kelurahan(id);


--
-- Name: kelurahan kelurahan_kecamatan_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.kelurahan
    ADD CONSTRAINT kelurahan_kecamatan_id_fkey FOREIGN KEY (kecamatan_id) REFERENCES public.kecamatan(id) ON DELETE CASCADE;


--
-- Name: pedagang_lokal pedagang_lokal_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pedagang_lokal
    ADD CONSTRAINT pedagang_lokal_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id) ON DELETE CASCADE;


--
-- Name: pengeluaran_harian pengeluaran_harian_alokasi_harian_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengeluaran_harian
    ADD CONSTRAINT pengeluaran_harian_alokasi_harian_id_fkey FOREIGN KEY (alokasi_harian_id) REFERENCES public.alokasi_harian(id) ON DELETE CASCADE;


--
-- Name: pengeluaran_harian pengeluaran_harian_pedagang_lokal_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengeluaran_harian
    ADD CONSTRAINT pengeluaran_harian_pedagang_lokal_id_fkey FOREIGN KEY (pedagang_lokal_id) REFERENCES public.pedagang_lokal(id) ON DELETE SET NULL;


--
-- Name: pengiriman pengiriman_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengiriman
    ADD CONSTRAINT pengiriman_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES public.drivers(id) ON DELETE RESTRICT;


--
-- Name: pengiriman pengiriman_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.pengiriman
    ADD CONSTRAINT pengiriman_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id) ON DELETE RESTRICT;


--
-- Name: posyandu posyandu_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.posyandu
    ADD CONSTRAINT posyandu_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id) ON DELETE RESTRICT;


--
-- Name: produksi_harian produksi_harian_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.produksi_harian
    ADD CONSTRAINT produksi_harian_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id);


--
-- Name: sekolah sekolah_sppg_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sekolah
    ADD CONSTRAINT sekolah_sppg_id_fkey FOREIGN KEY (sppg_id) REFERENCES public.sppg(id) ON DELETE SET NULL;


--
-- Name: sppg sppg_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.sppg
    ADD CONSTRAINT sppg_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: tokens tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: tracking tracking_pengiriman_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.tracking
    ADD CONSTRAINT tracking_pengiriman_id_fkey FOREIGN KEY (pengiriman_id) REFERENCES public.pengiriman(id) ON DELETE CASCADE;


--
-- Name: users_permissions users_permissions_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: users_permissions users_permissions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users_permissions
    ADD CONSTRAINT users_permissions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users users_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rian
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO rian;


--
-- PostgreSQL database dump complete
--

\unrestrict eXcn5F6zswsO6Wu3evG9gIrBUO7Ffg2LiIc9UdLQUQ5qfXEayZtO7c5G2LkHRlr

