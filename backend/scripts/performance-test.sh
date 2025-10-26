#!/bin/bash

# Performance test script for login endpoint
# Requires Apache Bench (ab) to be installed

echo "🚀 Starting Login Performance Test"
echo "=================================="

# Test configuration
URL="http://localhost:4000/api/auth/login"
CONCURRENT_USERS=10
TOTAL_REQUESTS=100
TEST_FILE="login_payload.json"

# Create test payload
cat > $TEST_FILE << EOF
{
    "username": "testuser",
    "password": "testpassword123"
}
EOF

echo "📊 Test Configuration:"
echo "  URL: $URL"
echo "  Concurrent Users: $CONCURRENT_USERS"
echo "  Total Requests: $TOTAL_REQUESTS"
echo ""

# Check if server is running
if ! curl -s "$URL" > /dev/null; then
    echo "❌ Server is not running at $URL"
    echo "Please start your server first: go run ./cmd/web"
    exit 1
fi

echo "✅ Server is running"
echo ""

# Run the performance test
echo "🔥 Running Apache Benchmark..."
echo "================================"
ab -n $TOTAL_REQUESTS -c $CONCURRENT_USERS -p $TEST_FILE -T application/json $URL

echo ""
echo "📈 Performance Analysis:"
echo "========================"
echo "Target: <50ms average response time"
echo ""
echo "🔍 To monitor slow requests in real-time:"
echo "tail -f your_log_file | grep 'SLOW REQUEST'"
echo ""
echo "🧹 Clean up test file:"
rm -f $TEST_FILE

echo ""
echo "✨ Test completed!"