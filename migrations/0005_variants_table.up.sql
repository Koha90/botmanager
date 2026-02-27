CREATE TABLE IF NOT EXISTS product_variants(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  district_id BIGINT REFERENCES districts(id) ON DELETE RESTRICT,
  pack_size TEXT NOT NULL,
  price BIGINT NOT NULL,
  archived_at TIMESTAMP NULL
);

CREATE INDEX idx_variants_product_id ON product_variants(product_id);
CREATE INDEX idx_variants_district_id ON product_variants(district_id);
