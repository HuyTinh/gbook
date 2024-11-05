<script lang="ts">
	let { name, price } = $props();
</script>

<div class="card w-80 shadow-md">
	<div class="flex justify-center">
		<figure class="w-1/2 py-5 text-center">
			<img
				src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
				alt="book"
			/>
		</figure>
	</div>
	<div class="card-body">
		<h2 class="card-title">{name}</h2>
		<p>{price}</p>
	</div>
</div>
