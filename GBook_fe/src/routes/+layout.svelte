<script lang="ts">
  import "../app.css";
  import Navbar from "../components/Navbar.svelte";
</script>

<div class="h-screen">
  <Navbar />
  <slot />
</div>
