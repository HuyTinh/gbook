<script lang="ts">
	import type { IBook } from '../api/books/book.type';
	import type { APIResponse } from '../types/api-response';
	import Card from './card.svelte';

	let { promise }: { promise: Promise<APIResponse<IBook[]>> } = $props();
</script>

{#await promise then response}
	{#each response.data as { name, price }}
		<Card {name} {price} />
	{/each}
{/await}
