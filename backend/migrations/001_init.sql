-- departments
create table if not exists departments(
  id serial primary key,
  name text unique not null
);

-- users (app-local mirror of Keycloak users)
create table if not exists users(
  id uuid primary key default gen_random_uuid(),
  keycloak_id text unique not null,
  email text not null,
  full_name text not null,
  department_id int references departments(id),
  role text not null check (role in ('user','manager','admin'))
);

-- vacations
create table if not exists vacations(
  id uuid primary key default gen_random_uuid(),
  user_id uuid references users(id) not null,
  start_date date not null,
  end_date date not null,
  total_days int not null,
  status text not null check (status in ('pending','approved','rejected')) default 'pending',
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

-- approvals
create table if not exists approvals(
  id uuid primary key default gen_random_uuid(),
  vacation_id uuid references vacations(id) not null,
  manager_id uuid references users(id) not null,
  decision text not null check (decision in ('approved','rejected')),
  comment text,
  decided_at timestamptz not null default now()
);

-- audit log
create table if not exists audit_logs(
  id bigserial primary key,
  actor_keycloak_id text not null,
  action text not null,
  entity text not null,
  entity_id text not null,
  payload jsonb,
  created_at timestamptz not null default now()
);

-- public holidays
create table if not exists public_holidays(
  id bigserial primary key,
  date date not null,
  local_name text not null,
  country_code text not null,
  unique(date,country_code)
);

create index if not exists idx_vacations_user on vacations(user_id);
