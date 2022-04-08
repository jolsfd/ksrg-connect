<script>
    export let result = null;
</script>

{#if result}
  {#if result.success}
    <div class="alert alert-success" role="alert">
      {result.msg}
    </div>
  {:else}
    <div class="alert alert-danger" role="alert">
      {result.msg}
    </div>
  {/if}
{/if}