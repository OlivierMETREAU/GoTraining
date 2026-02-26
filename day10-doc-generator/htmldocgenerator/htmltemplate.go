package htmldocgenerator

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go Project Documentation</title>
<style>
  body { margin: 0; font-family: system-ui, sans-serif; display: flex; height: 100vh; overflow: hidden; }
  #sidebar { width: 280px; background: #f5f5f5; border-right: 1px solid #ddd; padding: 10px; overflow-y: auto; box-sizing: border-box; }
  #content { flex: 1; padding: 20px; overflow-y: auto; }
  #searchBox { width: 100%; padding: 6px 8px; margin-bottom: 12px; font-size: 14px; border: 1px solid #ccc; border-radius: 4px; box-sizing: border-box; }
  .file-item { cursor: pointer; padding: 4px 8px; border-radius: 4px; font-size: 14px; }
  .file-item:hover { background: #e0e0e0; }
  .file-item.active { background: #1976d2; color: white; }
  .file-path { font-family: monospace; font-size: 12px; color: #555; }
  .package-name { font-weight: bold; font-size: 18px; margin-bottom: 4px; }
  .comment-block { margin-bottom: 16px; padding: 10px; border-left: 3px solid #1976d2; background: #fafafa; }
  .comment-context { font-weight: bold; font-size: 14px; color: #1976d2; margin-bottom: 6px; }
  .comment-text { white-space: pre-wrap; font-size: 14px; }
  .comment-lines { font-size: 11px; color: #888; margin-top: 4px; }
  mark { background: yellow; padding: 0 2px; }
  h3 { margin-top: 24px; margin-bottom: 8px; }
</style>
</head>

<body>

<div id="sidebar">
  <input id="searchBox" type="text" placeholder="Search files or content..." />

  <h3>Files</h3>
  {{range .Files}}
    <div class="file-item" 
         data-target="{{.ID}}" 
         data-name="{{.RelPath}}"
         data-content="{{range $ctx, $blocks := .GroupedDocs}}{{range $blocks}}{{.Text}} {{end}}{{end}}">
      <div>{{.RelPath}}</div>
      <div class="file-path">{{.Package}}</div>
    </div>
  {{end}}
</div>

<div id="content">
  {{range .Files}}
  <div class="file-doc" id="{{.ID}}" style="display:none;">
    <div class="package-name">package {{.Package}}</div>
    <div class="file-path">{{.RelPath}}</div>
    <hr>

    {{range $ctx, $blocks := .GroupedDocs}}
      <div class="context-group" data-context="{{$ctx}}">
        <h3>{{$ctx}}</h3>

        {{range $blocks}}
        <div class="comment-block" data-text="{{.Text}}">
          <div class="comment-context">
            {{$ctx}}{{if .SubContext}} — {{.SubContext}}{{end}}
          </div>
          <div class="comment-text">{{.Text}}</div>
          <div class="comment-lines">Lines {{.LineStart}}–{{.LineEnd}}</div>
        </div>
        {{end}}
      </div>
    {{end}}
  </div>
  {{end}}
</div>

<script>
(function() {
  const items = document.querySelectorAll('.file-item');
  const docs  = document.querySelectorAll('.file-doc');
  const searchBox = document.getElementById('searchBox');

  function show(id) {
    docs.forEach(d => d.style.display = (d.id === id ? 'block' : 'none'));
    items.forEach(i => i.classList.toggle('active', i.getAttribute('data-target') === id));
  }

  if (items.length > 0) {
    show(items[0].getAttribute('data-target'));
  }

  items.forEach(item => {
    item.addEventListener('click', () => {
      show(item.getAttribute('data-target'));
      applySearch(); // re-filter content when switching files
    });
  });

  function highlight(text, query) {
    if (!query) return text;
    const regex = new RegExp("(" + query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&') + ")", "gi");
    return text.replace(regex, "<mark>$1</mark>");
  }

  function applySearch() {
    const q = searchBox.value.toLowerCase();

    // Filter left panel (file list)
    items.forEach(item => {
      const name = item.getAttribute('data-name').toLowerCase();
      const content = item.getAttribute('data-content').toLowerCase();
      const match = name.includes(q) || content.includes(q);
      item.style.display = match ? 'block' : 'none';
    });

    // Filter right panel (content)
    docs.forEach(doc => {
      if (doc.style.display === 'none') return;

      const groups = doc.querySelectorAll('.context-group');

      groups.forEach(group => {
        let groupVisible = false;
        const blocks = group.querySelectorAll('.comment-block');

        blocks.forEach(block => {
          const text = block.getAttribute('data-text');
          const match = text.toLowerCase().includes(q);

          if (match) {
            block.style.display = 'block';
            block.querySelector('.comment-text').innerHTML = highlight(text, q);
            groupVisible = true;
          } else {
            block.style.display = 'none';
          }
        });

        group.style.display = groupVisible ? 'block' : 'none';
      });
    });
  }

  searchBox.addEventListener('input', applySearch);
})();
</script>

</body>
</html>
`
