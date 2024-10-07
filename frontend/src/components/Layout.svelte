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

<div class="p-2.5">
  <div>
    <nav class="flex justify-between items-center whitespace-nowrap text-center">
      <h1 class="inline-block mr-5"><a href="/">¿Who Knows?</a></h1>
      <div class="flex space-x-4">
        {#if $session.authenticated}
          <span>Logged in as {$session.username}</span>
          <button on:click={handleLogout}>Logout</button>
        {:else}
          <a class="float-right" href="/login">Log in</a>
          <a class="float-right" href="/register">Register</a>
        {/if}
      </div>
    </nav>
  </div>
  <div class="p-2.5">
    <slot></slot>
  </div>
  <div class="bg-gray-200 text-gray-500 p-1.5 text-xs">
    <span>¿Who Knows? &copy; 2009</span>
    <a href="/about">About</a>
  </div>
</div>
