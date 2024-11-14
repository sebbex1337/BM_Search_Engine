<script lang="ts">
  import { onMount } from "svelte";
  import { logout } from "../utils/api";
  import { session } from "../stores/session";
  import { navigate } from "svelte-routing";

  onMount(async () => {
    await session.checkSession();
  });

  async function handleLogout() {
    const res = await logout();
    if (res && res.ok) {
      session.checkSession();
      navigate("/");
    }
  }
</script>

<!-- Top-level container with Flexbox layout -->
<div class="min-h-screen flex flex-col p-2.5">
  
  <!-- Navigation Bar -->
  <nav class="flex justify-between items-center whitespace-nowrap text-center mb-4">
    <div class="flex items-center space-x-4">
      <h1 class="inline-block mr-5"><a href="/">¿Who Knows?</a></h1>
      <a href="/weather">Weather</a>
    </div>
    <div class="flex space-x-4">
      {#if $session.authenticated}
        <span>Logged in as {$session.username}</span>
        <button on:click={handleLogout} class="px-3 py-1 bg-blue-500 text-white rounded">Logout</button>
      {:else}
        <a href="/login" class="px-3 py-1 bg-green-500 text-white rounded">Log in</a>
        <a href="/register" class="px-3 py-1 bg-indigo-500 text-white rounded">Register</a>
      {/if}
    </div>
  </nav>

  <!-- Main Content Area -->
  <main class="flex-grow p-2.5">
    <slot></slot>
  </main>

  <!-- Footer -->
  <footer class="bg-gray-200 text-gray-700 p-4 text-xs flex flex-col sm:flex-row justify-between items-center">
    <span>¿Who Knows? &copy; 2009</span>
    <a href="/about" class="text-blue-500 hover:underline mt-2 sm:mt-0">About</a>
  </footer>

</div>