
use workApplication;
# 使用者
create table user(
	userId bigint not null AUTO_INCREMENT unique,
	companyCode varchar(50),
	userName varchar(20),
	employeeNumber varchar(30),
	account varchar(50) primary key,
	password varchar(50),
	onWorkDay timestamp,
	banch bigint,
	permession int,
	createTime timestamp,
	lastModify timestamp,
	monthSalary int,
	partTimeSalary int
);
alter table user auto_increment=1;

# 使用者偏好
create table userPreference(
	userId bigint primary key,
	style varchar(50),
	fontSize varchar(3),
	selfPhoto blob,
	createTime timestamp,
	lastModify timestamp
);
alter table userPreference add foreign key(userId) references user(userId) on update cascade on delete cascade;

# 班表
create table shift(
	shiftId bigint not null AUTO_INCREMENT unique,
	userId bigint,
	banchStyleId bigint,
	onShiftTime timestamp,
	offShiftTime timestamp,
	restTime time,
	punchIn timestamp,
	punchOut timestamp,
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
);
alter table shift add primary key (onShiftTime, offShiftTime, userId, restTime);
alter table shift add foreign key(userId) references user(userId) on update cascade on delete cascade;
alter table shift auto_increment=1;

# 換班
create table shiftChange(
	caseId bigint not null unique auto_increment,
	initiatorShiftId bigint,
	requestedShiftId bigint,
	reason varchar(200),
	caseProcess varchar(10),
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
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
	reason varchar(200),
	caseProcess varchar(10),
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
);
alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);
alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 忘記打卡
create table forgetPunch(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	targetPunch varchar(3),
	reason varchar(200),
	caseProcess varchar(10),
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
);
alter table forgetPunch add primary key (shiftId, targetPunch);
alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 請假
create table dayOff(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	dayOffType varchar(10),
	reason varchar(200),
	caseProcess varchar(10),
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
);
alter table dayOff add primary key (shiftId, dayOffType);
alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 遲到早退
create table lateExcused(
	caseId bigint not null unique auto_increment,
	shiftId bigint,
	lateExcusedType varchar(10),
	reason varchar(200),
	caseProcess varchar(10),
	specifyTag varchar(50),
	createTime timestamp,
	lastModify timestamp
);
alter table lateExcused add primary key (shiftId, lateExcusedType);
alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;

# 公司
create table company(
	companyId bigint not null unique auto_increment,
	companyCode varchar(50) unique,
	companyName varchar(200),
	companyLocation varchar(200),
	companyPhoneNumber varchar(20),
	bossId bigint not null,
	settlementDate int,
	termStart timestamp,
	termEnd timestamp,
	createTime timestamp,
	lastModify timestamp
);
alter table company add primary key (companyId, companyCode);
alter table company auto_increment=1;

# 公司部門
create table companyBanch(
	id bigint not null unique auto_increment,
	companyId bigint,
	banchName varchar(50),
	banchShiftStyle varchar(200),
	createTime timestamp,
	lastModify timestamp
);
alter table companyBanch add primary key (companyId, banchName);
alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade on delete cascade;

# 部門圖標
create table banchStyle(
	styleId bigint not null unique auto_increment,
	banchId bigint,
	icon varchar(100),
	restTime time,
	timeRangeName varchar(20),
	onShiftTime time,
	offShiftTime time,
	createTime timestamp,
	lastModify timestamp
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
	userName varchar(20),
	employeeNumber varchar(30),
	account varchar(50),
	onWorkDay timestamp,
	banch bigint,
	permession int,
	monthSalary int,
	partTimeSalary int,
	createTime timestamp,
	lastModify timestamp
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
	specifyTag varchar(50),
	isAccept int,
	createTime timestamp,
	lastModify timestamp
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
create table workTIme(
	workTimeId bigInt not null unique auto_increment,
	userId bigInt,
	year int,
	month int,
	workHours int,
	timeOff int,
	usePaidVocation int,
	createTime timestamp,
	lastModify timestamp
);
alter table workTime add primary key(`userId`, `year`, `month`);
alter table workTime add foreign key(`userId`) references user(`userId`) on update cascade on delete cascade;
alter table workTime auto_increment=1;

# 個人特休
create table paidVocation(
	paidVocationId bigInt not null unique auto_increment,
	userId bigInt,
	year int,
	count int,
	createTime timestamp,
	lastModify timestamp
);
alter table paidVocation add primary key(`userId`, `year`);
alter table paidVocation add foreign key(`userId`) references user(`userId`) on update cascade on delete cascade;
alter table paidVocation auto_increment=1;

 
