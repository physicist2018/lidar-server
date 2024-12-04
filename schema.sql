CREATE TABLE Experiment(
    ID integer primary key autoincrement,
    Start_Time timestamp not null default now,
    Title varchar(100) not null default 'no-title',
    Comments varchar(500) not null default 'no-comments',
    Wavelen real not null default 532.0,
    Vert_Res real not null default 1500.0,
    Accum integer not null default 10,
    unique (Start_Time, Title, Accum)
);

CREATE TABLE Measurement(
    ID integer primary key autoincrement,
    Start_Time timestamp not null default now,
    Prof_Cnt integer not null default 1,
    Prof_Len integer not null default 512,
    Rep_Rate integer not null default 10,
    Profile_Data json,
    Experiment_ID integer,
    CONSTRAINT experiment_fk FOREIGN KEY (Experiment_ID) REFERENCES Experiment(ID)
);
