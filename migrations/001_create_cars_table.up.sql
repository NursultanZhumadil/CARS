CREATE TABLE cars (
                      id SERIAL PRIMARY KEY,
                      brand VARCHAR(100) NOT NULL,
                      model VARCHAR(100) NOT NULL,
                      year INT NOT NULL,
                      price NUMERIC NOT NULL,
                      mileage INT NOT NULL,
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                      updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
