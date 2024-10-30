<script lang="ts">
  import { onMount } from "svelte";
  import { writable } from "svelte/store";
  import Layout from "../components/Layout.svelte";
  import type { Page } from "../utils/types";
  import { search } from "../utils/api";

  let searchInput: HTMLInputElement;
  let query = writable("");
  let searchResults = writable<Page[]>([]);

  onMount(() => {
    searchInput.focus();
  });

  async function makeSearchRequest() {
    const data = await search($query)
    searchResults.set(data.data);
  }
</script>

<Layout>
  <div class="flex items-center space-x-2 p-4">
    <input
      bind:this={searchInput}
      placeholder="Search..."
      bind:value={$query}
      on:keypress={(event) => {
        if (event.key === "Enter") makeSearchRequest();
      }}
      class="flex-grow p-2 border border-gray-300 rounded"
    />
    <button on:click={makeSearchRequest} class="py-2 px-8 border border-blue-600 rounded bg-blue-600 cursor-pointer"
      >Search</button
    >
  </div>

  <div class="p-4">
    {#if $searchResults.length > 0}
      <h2 class="text-xl font-semibold mb-4">Results</h2>

      {#each $searchResults as page}
        <div class="mb-4 bg-gray-200 rounded p-2">
          <h3 class="text-lg font-medium"><a href={page.URL} class="text-blue-500">{page.title}</a></h3>
        </div>
      {/each}
    {/if}
  </div>
</Layout>
