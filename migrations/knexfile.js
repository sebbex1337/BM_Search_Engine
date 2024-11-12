import "dotenv/config";

/**
 * @type { import("knex").Knex.Config }
 */
export default {
  postgres: {
    client: "postgresql",
    connection: {
      database: process.env.POSTGRES_DB,
      user: process.env.POSTGRES_USER,
      password: process.env.POSTGRES_PASSWORD,
      host: process.env.POSTGRES_HOST,
      port: process.env.POSTGRES_PORT,
    },
  },
  sqlite: {
    client: "sqlite3",
    connection: {
      filename: "../backend/data/whoknows.db",
    },
  },
};
