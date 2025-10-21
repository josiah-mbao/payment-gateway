# payment-gateway

# 🧃 Juice Coin - PesaFlow Sandbox

> **Experience the future of African cross-border payments with virtual Juice Coins. Real technology, zero risk.**

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![React](https://img.shields.io/badge/Next.js-14.0+-black.svg)](https://nextjs.org)
[![Interledger](https://img.shields.io/badge/Interledger-Rafiki-green.svg)](https://interledger.org)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 🎯 What is Juice Coin?

Juice Coin is a **virtual currency demonstration** of PesaFlow - a cross-border payment platform designed specifically for African businesses. It showcases how modern payment technology can solve Africa's cross-border payment challenges using the Interledger Protocol.

### 🌍 The Problem
- African SMEs struggle to accept payments from other African countries
- Traditional cross-border payments take 3-5 days with 8-15% fees
- No simple APIs exist for developers to build cross-border features
- Fragmented financial systems across African countries

### 💡 Our Solution
- **Virtual Economy**: Experience payments with Juice Coins (JUC)
- **Real Technology**: Actual Interledger Protocol integration
- **Instant Settlement**: Cross-border payments in minutes, not days
- **Developer First**: Full API access for building integrations

## 🚀 Features

### 🎮 Virtual Economy
- **10,000 JUC** starting balance for all users
- **Virtual merchants** from Kenya, Nigeria, Ghana, South Africa
- **Demo products** and services to simulate real transactions
- **Cross-border payments** between virtual African countries

### 🔧 Real Technology Stack
- **Go Backend**: High-performance payment processing
- **Interledger Protocol**: Real Rafiki integration for cross-border settlements
- **React Dashboard**: Modern merchant interface
- **PostgreSQL**: Robust data storage
- **Kubernetes**: Production-ready deployment

### 🛡️ Safety & Compliance
- ✅ **No real money** involved
- ✅ **Virtual currency only** - JUC has no monetary value
- ✅ **Educational purpose** - Technology demonstration
- ✅ **Regulatory safe** - No financial services licensing required

## 🏗️ Architecture

```mermaid
graph TB
    A[Frontend Dashboard] --> B[Go API Server]
    B --> C[Juice Coin Engine]
    B --> D[Interledger Rafiki]
    C --> E[Virtual Ledger]
    D --> F[Cross-border Settlement]
    B --> G[PostgreSQL]
    H[Webhooks] --> B
    
    style C fill:#e1f5fe
    style D fill:#f3e5f5

## 🛠️ Tech Stack

### Backend
- **Go 1.21+** - High-performance API server
- **Gin Framework** - HTTP web framework
- **GORM** - Database ORM
- **PostgreSQL** - Primary database
- **Rafiki Money** - Interledger connector

### Frontend
- **Next.js 14** - React framework
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **Chart.js** - Analytics and metrics

### Infrastructure
- **Docker** - Containerization
- **Kubernetes** - Orchestration
- **Prometheus** - Monitoring
- **Grafana** - Dashboards

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- Node.js 18+
- Docker (optional)
