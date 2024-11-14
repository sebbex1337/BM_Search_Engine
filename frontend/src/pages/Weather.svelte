<script lang="ts">
  import { writable } from "svelte/store";
  import Layout from "../components/Layout.svelte";
  import { onMount } from "svelte";
  import { getWeather } from "../utils/api";
  import type { weather } from "../utils/types";

  // Create a writable store for the weather data
  let weatherData = writable<weather | null>(null);
  onMount(async () => {
    const response = await getWeather();
    weatherData.set(response);
  });
</script>

<Layout>
  <h1>Weather</h1>
  <br />
  {#if $weatherData}
    <h1>City: {$weatherData.name}</h1>
    <h1>Temperature: {$weatherData.main.temp} °C</h1>
  {:else}
    <h1>Loading...</h1>
  {/if}
</Layout>
