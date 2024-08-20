CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    dob DATE,
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(50),
    role VARCHAR(20) DEFAULT 'customer',
    createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updateAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Rooms (
    room_id SERIAL PRIMARY KEY,
    room_type VARCHAR(20),
    description TEXT,
    price INT NOT NULL,
    room_status BOOLEAN,
    max_guest INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Bookings (
    booking_id SERIAL PRIMARY KEY,
    room_id INT,
    user_id INT,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    booking_status VARCHAR(50) DEFAULT 'Confirmed',
    FOREIGN KEY (room_id) REFERENCES Rooms(room_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Room_Images (
    image_id SERIAL PRIMARY KEY,
    room_id INT,
    image_url VARCHAR(255) NOT NULL,
    FOREIGN KEY (room_id) REFERENCES Rooms(room_id)
);

CREATE TABLE IF NOT EXISTS Reviews (
    review_id SERIAL PRIMARY KEY,
    room_id INT,
    user_id INT,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    review_text TEXT,
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES Rooms(room_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Payments (
    payment_id SERIAL PRIMARY KEY,
    booking_id INT,
    payment_date DATE NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method VARCHAR(50),
    payment_status VARCHAR(50) DEFAULT 'Paid',
    FOREIGN KEY (booking_id) REFERENCES Bookings(booking_id)
);
