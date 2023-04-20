
use workApplication;
# 使用者
create table user(
	userId bigint not null AUTO_INCREMENT unique,
	companyCode varchar(50) default '',
	userName varchar(20) default '',
	employeeNumber varchar(30) default '',
	account varchar(50) primary key,
	password varchar(50) default '',
	onWorkDay timestamp default now(),
	banch bigint default -1,
	permession int default 2,
	createTime timestamp default now(),
	lastModify timestamp default now(),
	monthSalary int default 0,
	partTimeSalary int default 0
);
alter table user auto_increment=1;

# 使用者偏好
create table userPreference(
	userId bigint primary key,
	style varchar(50) default '',
	fontSize varchar(3) default '',
	selfPhoto blob default null,
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table userPreference add foreign key(userId) references user(userId) on update cascade on delete cascade;

# 班表
create table shift(
	shiftId bigint not null AUTO_INCREMENT unique,
	userId bigint,
	banchStyleId bigint default -1,
	banchId bigInt default -1,
	year int default -1,
	month int default -1,
	icon string varchar(100) default '',
	onShiftTime timestamp,
	offShiftTime timestamp,
	restTime time,
	punchIn timestamp default null,
	punchOut timestamp default null,
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table shift add primary key (onShiftTime, offShiftTime, userId, restTime);
alter table shift add foreign key(userId) references user(userId) on update cascade on delete cascade;
alter table shift auto_increment=1;

# 換班
create table shiftChange(
	caseId bigint not null unique auto_increment,
	initiatorShiftId bigint,
	requestedShiftId bigint,
	reason varchar(200) default '',
	caseProcess varchar(10) default '',# ok, wait, manageCheck, reject
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table shiftChange add primary key (initiatorShiftId, requestedShiftId);
alter table shiftChange add foreign key(initiatorShiftId) references shift(shiftId) on update cascade on delete cascade;
alter table shiftChange add foreign key (requestedShiftId) references shift(shiftId) on update cascade on delete cascade;

# 加班
create table shiftOverTime(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	initiatorOnOverTime timestamp,
	initiatorOffOverTime timestamp,
	reason varchar(200) default '',
	caseProcess varchar(10) default '',
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);
alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 忘記打卡
create table forgetPunch(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	targetPunch varchar(3), # on off
	reason varchar(200) default '',
	caseProcess varchar(10) default '',
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table forgetPunch add primary key (shiftId, targetPunch);
alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 請假
create table dayOff(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	dayOffType varchar(10),
	reason varchar(200) default '',
	caseProcess varchar(10) default '',
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table dayOff add primary key (shiftId, dayOffType);
alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 遲到早退
create table lateExcused(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	lateExcusedType varchar(10),
	reason varchar(200) default '',
	caseProcess varchar(10) default '',
	specifyTag varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table lateExcused add primary key (shiftId, lateExcusedType);
alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 公司
create table company(
	companyId bigint not null unique auto_increment,
	companyCode varchar(50) unique,
	companyName varchar(200) default '',
	companyLocation varchar(200) default '',
	companyPhoneNumber varchar(20) default '',
	bossId bigint not null default -1,
	settlementDate int default -1,
	termStart timestamp default now(),
	termEnd timestamp default now(),
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table company add primary key (companyId, companyCode);
alter table company auto_increment=1;

# 公司部門
create table companyBanch(
	id bigint not null unique auto_increment,
	companyId bigint,
	banchName varchar(50),
	banchShiftStyle varchar(200) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table companyBanch add primary key (companyId, banchName);
alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade on delete cascade;

# 部門圖標
create table banchStyle(
	styleId bigint not null unique auto_increment,
	banchId bigint,
	icon varchar(100),
	restTime time,
	timeRangeName varchar(20) default '',
	onShiftTime time,
	offShiftTime time,
	delFlag varchar(1) default 'N'
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table banchStyle auto_increment=1;
alter table banchStyle add primary key(banchId, restTime, onShiftTime, offShiftTime, icon);
alter table banchStyle add foreign key (banchId) references companyBanch(id) on update cascade on delete cascade;

# 部門規則
create table banchRule(
	ruleId bigint not null unique auto_increment,
	banchId bigint,
	maxPeople int,
	minPeople int,
	weekDay int,
	weekType int,
	onShiftTime time,
	offShiftTime time,
	createTime timestamp,
	lastModify timestamp
);
alter table banchRule auto_increment=1;
alter table banchRule add primary key (banchId, weekDay, weekType);
alter table banchRule add foreign key (banchId) references companyBanch(id) on update cascade on delete cascade;

# 離職員工
create table quitWorkUser(
	quitId bigint not null unique auto_increment,
	userId bigint,
	companyCode varchar(20),
	userName varchar(20) default '',
	employeeNumber varchar(30) default '',
	account varchar(50) default '',
	onWorkDay timestamp default now(),
	banch bigint default -1,
	permession int default 2,
	monthSalary int default 0,
	partTimeSalary int default 0,
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table quitWorkUser auto_increment=1;
alter table quitWorkUser add primary key (userId, companyCode);
alter table quitWorkUser add foreign key (userId) references user(userId) on update cascade on delete cascade;
alter table quitWorkUser add foreign key (companyCode) references company(companyCode) on update cascade on delete cascade;

# 等待公司回覆
create table waitCompanyReply(
	waitId bigint not null unique auto_increment,
	userId bigint,
	companyId bigint,
	specifyTag varchar(50) default '',
	isAccept int default 1,
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table waitCompanyReply auto_increment=1;
alter table waitCompanyReply add primary key (userId, companyId);
alter table waitCompanyReply add foreign key (userId) references user(userId) on update cascade on delete cascade;
alter table waitCompanyReply add foreign key (companyId) references company(companyId) on update cascade on delete cascade;

# 假日設定
create table weekendSetting(
	weekendId bigint not null unique auto_increment,
	companyId bigint,
	date date,
	createTime timestamp,
	lastModify timestamp
);
alter table weekendSetting auto_increment=1;
alter table weekendSetting add primary key (date, companyId);
alter table weekendSetting add foreign key (companyId) references company(companyId) on update cascade on delete cascade;

# 個人時數
create table workTime(
	workTimeId bigInt not null unique auto_increment,
	userId bigInt,
	year int,
	month int,
	workHours int default 8,
	timeOff int default 0,
	usePaidVocation int default 0,
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table workTime add primary key(`userId`, `year`, `month`);
alter table workTime add foreign key(`userId`) references user(`userId`) on update cascade on delete cascade;
alter table workTime auto_increment=1;

# 個人特休
create table paidVocation(
	paidVocationId bigInt not null unique auto_increment,
	userId bigInt,
	year int,
	count int default 0,
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table paidVocation add primary key(`userId`, `year`);
alter table paidVocation add foreign key(`userId`) references user(`userId`) on update cascade on delete cascade;
alter table paidVocation auto_increment=1;

# 績效評核
create table performance(
	performanceId bigint not null unique auto_increment,
	userId bigInt not null,
	year int,
	month int,
	banchId bigInt default -1,
	goal varchar(1000) default '',
	attitude int default 0,
	efficiency int default 0,
	professional int default 0,
	directions varchar(1000) default '',
	beLate int default 0,
	dayOffNotOnRule int default 0,
	banchName varchar(50) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);
alter table performance add primary key(`userId`, `year`, `month`);
alter table performance add foreign key(`userId`) references user(`userId`) on update cascade on delete cascade;

# 紀錄
create table log(
	logId bigint not null unique auto_increment,
	userId bigInt default -1,
	userName varchar(50) default '',
	companyId bigInt default -1,
	companyCode varchar(50) default '',
	permession int default 2,
	routes varchar(100) default '',
	ip varchar(100) default '',
	params varchar(1000) default '',
	msg varchar(1000) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);

# 編輯 班表歷程
create table shiftEditLog (
	logId bigint not null unique auto_increment,
	year int,
	month int,
	banchId bigInt default -1,
	msg varchar(1000) default '',
	createTime timestamp default now(),
	lastModify timestamp default now()
);