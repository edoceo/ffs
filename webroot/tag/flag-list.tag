<flag-list>

<table class="table">
<thead>
	<tr>
		<th>Flag</th>
		<th>Enum</th>
		<th>MRU</th>
		<th>Rollout</th>
		<th>BL</th>
		<th>WL</th>
	</tr>
</thead>
<tbody>
	<tr each={ item in flag_list }>
		<td><a href="/flag.html?id={ item.id }">{ item.stub }</a></td>
		<td>{ item.name }</td>
		<td>MMM/NN</td>
		<td>{item.green_list}</td>
		<td>{item.white_list}</td>
		<td>{item.black_list}</td>
	</tr>
</tbody>
</table>

<script>
var tag = this;

tag.on('mount', function() {

	$.get('/api/v2016/flags', function(res) {
		tag.update({ flag_list: res.Flag });
	});

});
</script>

</flag-list>
