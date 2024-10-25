<script>
  // @ts-nocheck
  import { Badge, Button, Card, Group, Image, Text } from "@svelteuidev/core";

  export let title;
  export let price;
</script>

<Card shadow="sm" padding="lg" class="w-80">
  <Card.Section first padding="lg">
    <Image
      src="https://images.unsplash.com/photo-1529158062015-cad636e205a0?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8dGhlJTIwZ3JlYXQlMjBnYXNieSUyMGJvb2t8ZW58MHx8MHx8fDA%3D"
      height={160}
      alt="Portugal"
    />
  </Card.Section>

  <Group position="apart">
    <Text weight={500}>{title}</Text>
    <Badge color="pink" variant="light">On Sale</Badge>
  </Group>

  <Text size="sm">
    {price}
  </Text>

  <Button variant="light" color="blue" fullSize>Book classic tour now</Button>
</Card>
