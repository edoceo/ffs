<user-list>

<table class="table">
<thead>
	<tr>
		<th>User</th>
		<th>Hash</th>
		<th>Seen</th>
	</tr>
</thead>
<tbody>
	<tr each={ item, i in user_list}>
		<td><a href="/user.html?u={ item.hash }">{ item.name }</a></td>
		<td>{ item.hash }</td>
		<td>{ item.access_at }</td>
	</tr>
</tbody>
</table>

<script>
var tag = this;

tag.on('mount', function() {

	$.get('/api/v2016/users', function(res) {
		tag.update({ user_list: res });
	});

});
</script>

</user-list>
