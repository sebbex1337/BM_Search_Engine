import knexConfig from "./knexfile.js";
import knex from "knex";

async function migrate() {
  const knexSQLite = knex(knexConfig.sqlite);
  const knexPostgres = knex(knexConfig.postgres);
  try {
    const pages = await knexSQLite.select("*").from("pages");
    await knexPostgres("pages").insert(pages);
    console.log("Migration successful");
  } catch (error) {
    console.error("Migration failed: ", error);
  } finally {
    await knexSQLite.destroy();
    await knexPostgres.destroy();
  }
}

migrate();
