<script type="module">
  import { Tabs } from "@svelteuidev/core";
  import { Camera, EnvelopeClosed, Gear } from "radix-icons-svelte";
  import Card from "../components/Card.svelte";
  import { BookService } from "../api/services/book.service";
  import { onMount } from "svelte";

  // @ts-ignore
  /**
   * @type {import("../types/book.type").IBook[]}
   */
  let books = [];
  onMount(async () => {
    books = await BookService.getAllBook();
  });
</script>

<div>
  <Tabs color="red" tabPadding="md">
    <Tabs.Tab label="Gallery" icon={Camera}>
      <div class="flex flex-wrap justify-evenly gap-2 p-10">
        {#each books as book}
          <Card title={book.title} price={1000} />
        {/each}
      </div>
    </Tabs.Tab>
    <Tabs.Tab label="Messages" icon={EnvelopeClosed}
      >Messages tab content</Tabs.Tab
    >
    <Tabs.Tab label="Settings" icon={Gear}>Settings tab content</Tabs.Tab>
  </Tabs>
</div>
