CREATE TABLE TypeRoom (
    type_id uuid PRIMARY KEY,
    type_name varchar(50) NOT NULL,
    description text, 
    price_per_night decimal(10, 2) NOT NULL,
    max_occupancy int NOT NULL,
    room_size decimal(5, 2),
    image_url varchar(255),
    status varchar(20) DEFAULT 'active',
    discount decimal(5, 2) DEFAULT 0
);

CREATE TABLE Room (
    room_id uuid PRIMARY KEY,
    room_name varchar(50) NOT NULL,
    type_id uuid REFERENCES TypeRoom(type_id),  
    floor int NOT NULL,
    status varchar(20) DEFAULT 'available',
    price_override decimal(10, 2),
    cleaning_status varchar(20) DEFAULT 'clean',
    check_in_time timestamp,
    check_out_time timestamp,
    current_guest varchar(100),
    note text
);

create table Customer (
	customer_id uuid primary key,
	full_name varchar (100) not null,
	email varchar(100),
	phone_number varchar(15) not null,
	address text,
	nationality varchar(50),
	date_of_birth date,
	id_document varchar(50),
	registration_date timestamp default current_timestamp,
	note text
);

create table Booking (
	booking_id uuid primary key,
	customer_id uuid references Customer(customer_id),
	room_id uuid references Room(room_id),
	booking_date timestamp default current_timestamp,
	check_in_date timestamp not null,
	check_out_date timestamp not null,
	total_price decimal (10, 2) not null,
	status varchar(20) default 'booked',
	payment_status varchar(20) default 'pending',
	note text
);

create table Payment(
	payment_id uuid primary key,
	booking_id uuid references Booking(booking_id),
	customer_id uuid references Customer(customer_id),
	payment_date timestamp default current_timestamp,
	amount decimal(10, 2) not null,
	payment_method varchar(50) not null,
	payment_status varchar(20) default 'pending',
	note text
);

create table Employee (
	employee_id uuid primary key,
	full_name varchar(100) not null,
	email varchar(100),
	phone_number varchar(15) not null,
	address text,
	position varchar(50) not null,
	salary decimal (10, 2) not null,
	hire_date date not null,
	date_of_birth date,
	id_document varchar(50),
	status varchar(20) default 'active',
	note text
);

create table Salary (
	salary_id uuid primary key,
	employee_id uuid references Employee(employee_id),
	base_salary decimal(10, 2) not null check(base_salary >= 0),
	bonus decimal(10, 2) default 0 check (bonus >= 0),
	allowance decimal(10, 2) default 0 check (allowance >= 0),
	deduction decimal (10, 2) default 0 check (deduction >= 0),
	net_salary decimal(10, 2) not null check(net_salary >= 0),
	payment_date date not null,
	note text
);

-- Add to the table of TypeRoom
ALTER TABLE TypeRoom 
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Room
ALTER TABLE Room
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Customer
alter table customer 
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Booking
alter table booking
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Payment
alter table payment 
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Employee
alter table employee 
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add to the table of Salary
alter table salary 
	ADD createTime timestamp,
	ADD createBy varchar(100),
	ADD updateTime timestamp,
	ADD updateBy varchar(100),
	ADD deleteTime timestamp,
	ADD deleteBy varchar(100);

-- Add eployee_id to the booking table
alter table booking 
add employee_id uuid references Employee(employee_id);

-- Display column eployee_id
alter table booking 
alter column employee_id set not null;

--Delete Customer_Id in Payment Table
alter table payment 
drop column customer_id;