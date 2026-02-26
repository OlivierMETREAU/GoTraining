package htmldocgenerator

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go Project Documentation</title>
<style>
  body {
    margin: 0;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    display: flex;
    height: 100vh;
    overflow: hidden;
  }
  #sidebar {
    width: 280px;
    background: #f5f5f5;
    border-right: 1px solid #ddd;
    padding: 10px;
    box-sizing: border-box;
    overflow-y: auto;
  }
  #content {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
  }
  .file-item {
    cursor: pointer;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 14px;
  }
  .file-item:hover {
    background: #e0e0e0;
  }
  .file-item.active {
    background: #1976d2;
    color: white;
  }
  .file-path {
    font-family: monospace;
    font-size: 12px;
    color: #555;
  }
  .package-name {
    font-weight: bold;
    font-size: 18px;
    margin-bottom: 4px;
  }
  .comment-block {
    margin-bottom: 16px;
    padding: 10px;
    border-left: 3px solid #1976d2;
    background: #fafafa;
  }
  .comment-context {
    font-weight: bold;
    font-size: 13px;
    color: #1976d2;
    margin-bottom: 4px;
  }
  .comment-text {
    white-space: pre-wrap;
    font-size: 14px;
  }
  .comment-lines {
    font-size: 11px;
    color: #888;
    margin-top: 4px;
  }
</style>
</head>
<body>
<div id="sidebar">
  <h3>Files</h3>
  {{range .Files}}
    <div class="file-item" data-target="{{.ID}}">
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
    {{if .Comments}}
      {{range .Comments}}
      <div class="comment-block">
        <div class="comment-context">{{.Context}}</div>
        <div class="comment-text">{{.Text}}</div>
        <div class="comment-lines">Lines {{.LineStart}}–{{.LineEnd}}</div>
      </div>
      {{end}}
    {{else}}
      <p>No comments found in this file.</p>
    {{end}}
  </div>
  {{end}}
</div>

<script>
  (function() {
    const items = document.querySelectorAll('.file-item');
    const docs  = document.querySelectorAll('.file-doc');

    function show(id) {
      docs.forEach(d => d.style.display = (d.id === id ? 'block' : 'none'));
      items.forEach(i => i.classList.toggle('active', i.getAttribute('data-target') === id));
    }

    if (items.length > 0) {
      const firstId = items[0].getAttribute('data-target');
      show(firstId);
    }

    items.forEach(item => {
      item.addEventListener('click', () => {
        const id = item.getAttribute('data-target');
        show(id);
      });
    });
  })();
</script>
</body>
</html>
`
