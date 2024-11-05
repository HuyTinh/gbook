<script lang="ts">
	import '../app.css';
	import { Icon } from 'svelte-icons-pack';
	import { BsCart2 } from 'svelte-icons-pack/bs';
</script>

<div class="navbar bg-blue-400">
	<div class="flex-1">
		<a href="/" class="btn btn-ghost text-xl text-white">GBook</a>
	</div>
	<div class="flex-none">
		<div class="dropdown dropdown-end">
			<div tabindex="0" role="button" class="btn btn-ghost btn-circle">
				<div class="indicator">
					<Icon src={BsCart2} size={24} />
					<span class="badge badge-sm indicator-item">8</span>
				</div>
			</div>
			<div class="card card-compact dropdown-content bg-base-100 z-[1] mt-3 w-52 shadow">
				<div class="card-body">
					<span class="text-lg font-bold">8 Items</span>
					<span class="text-info">Subtotal: $999</span>
					<div class="card-actions">
						<button class="btn btn-primary btn-block">View cart</button>
					</div>
				</div>
			</div>
		</div>
		<div class="dropdown dropdown-end">
			<div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
				<div class="w-10 rounded-full">
					<img
						alt="Tailwind CSS Navbar component"
						src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
					/>
				</div>
			</div>
			<ul class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
				<li>
					<a href="/" class="justify-between">
						Profile
						<span class="badge">New</span>
					</a>
				</li>
				<li>
					<a href="/">Settings</a>
				</li>
				<li><a href="/">Logout</a></li>
			</ul>
		</div>
	</div>
</div>

<slot />
