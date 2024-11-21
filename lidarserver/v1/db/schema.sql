CREATE TABLE Experiment(
    ID integer primary key autoincrement,
    starttime datetime not null default now,
    title varchar(100) not null default 'no-title',
    comments varchar(500) not null default 'no-comments',
    wavelen real not null default 532.0,
    vertres real not null default 1500.0,
    accum integer not null default 10,
    unique (starttime, title, accum)
);

CREATE TABLE Measurement(
    ID integer primary key autoincrement,
    starttime datetime not null default now,
    profcnt integer not null default 1,
    proflen integer not null default 512,
    reprate integer not null default 10,
    prof_type varchar(3) check (prof_type in ('dat','dak')) not null default 'dat',
    profile json,
    experiment_id integer,
    CONSTRAINT experiment_fk FOREIGN KEY (experiment_id) REFERENCES Experiment(id)
);

