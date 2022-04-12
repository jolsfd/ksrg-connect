<script>
    export let loading;
</script>

{#if loading}
  <div class="d-flex align-items-center">
    <strong>Loading...</strong>
    <div class="spinner-border ms-auto" role="status" aria-hidden="true" />
  </div>
{/if}