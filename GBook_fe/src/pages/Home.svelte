<script>
    import { Tabs } from "@svelteuidev/core";
    import { Camera, EnvelopeClosed, Gear } from "radix-icons-svelte";
</script>

<div>
    <Tabs color="red" tabPadding="md">
        <Tabs.Tab label="Gallery" icon={Camera}>Gallery tab content</Tabs.Tab>
        <Tabs.Tab label="Messages" icon={EnvelopeClosed}
            >Messages tab content</Tabs.Tab
        >
        <Tabs.Tab label="Settings" icon={Gear}>Settings tab content</Tabs.Tab>
    </Tabs>
</div>
