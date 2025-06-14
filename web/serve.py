#!/usr/bin/env python3
"""
Simple HTTP server for wherewasi HTMX demo
Serves static files and simulates the context API
"""

import http.server
import socketserver
import json
import os
from urllib.parse import urlparse, parse_qs

class WhereWasIHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        # Serve index.html for root request
        if self.path == '/':
            self.path = '/index.html'
        return super().do_GET()

    def do_POST(self):
        if self.path == '/api/pull-context':
            self.send_context_response()
        else:
            self.send_error(404, "Not Found")

    def send_context_response(self):
        # Simulate wherewasi context generation
        context_data = """🪂 WHEREWASI CONTEXT - Development Session
═══════════════════════════════════════

📍 PROJECT: wherewasi (AI context generation CLI)
📅 TIMEFRAME: Last 7 days
⚡ ACTIVE SESSION: HTMX demo development detected

🎯 PROJECT SUMMARY:
Local-first context generation CLI for AI collaboration. Working MVP with
ecosystem tracking, ripcord-style emergency context deployment, and cross-
project intelligence gathering.

📝 RECENT ACTIVITY:
• HTMX landing page created with ripcord theme
• Framework experimentation: Fresh/Deno for osmotic
• SlopSquid rebrand to CLI tool with tentacles concept
• Domain portfolio strategy for framework testing

🔍 KEY FILES:
main.go → Core CLI with SQLite context storage
web/index.html → HTMX demo with parachute animations
README.md → Usage examples and ecosystem integration
internal/database/ → Cross-project tracking logic

💬 RECENT CONVERSATIONS:
AI_SESSION.md:156-234 → Context generation architecture
cursor_htmx_demo.md:45-67 → Interactive ripcord implementation

🚀 NEXT PRIORITIES:
• Connect demo to real wherewasi CLI instance
• Add miqro.dev Astro site development
• Integrate context sharing across QRY ecosystem

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Context ready for AI collaboration → Copied to clipboard"""

        # Send HTML response for HTMX
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()

        # Return pre-formatted context for display
        response_html = f'<pre style="white-space: pre-wrap;">{context_data}</pre>'
        self.wfile.write(response_html.encode('utf-8'))

    def log_message(self, format, *args):
        # Custom logging
        print(f"🪂 {self.address_string()} - {format % args}")

def run_server(port=8080):
    # Change to the web directory
    web_dir = os.path.dirname(os.path.abspath(__file__))
    os.chdir(web_dir)

    with socketserver.TCPServer(("", port), WhereWasIHandler) as httpd:
        print(f"🪂 WhereWasI HTMX Demo Server")
        print(f"📡 Serving at http://localhost:{port}")
        print(f"📁 Directory: {web_dir}")
        print(f"🎯 Pull the ripcord to test context generation!")
        print("━" * 50)

        try:
            httpd.serve_forever()
        except KeyboardInterrupt:
            print("\n🪂 Ripcord pulled! Server landing safely...")

if __name__ == "__main__":
    run_server()
