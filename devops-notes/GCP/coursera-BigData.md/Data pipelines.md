# Data pipeline challenges
Real time streaming data
Data can come from different sources (IoT sensors, ...)
Batch and streaming data
# Workflow
1. Ingest: Capture data (sensors, ...)
2. Store & Distribute: Pub/Sub
3. Process: ETL into datawarehouse (Cloud Dataflow)
Apache Beam is a data processing programming model (SDK for writing data pipelines)
Data can come from different sources (PubSub/Storage)
Output can be: PubSub/BigQuery/Cloud Storage

4. Store & Analysis (BigQuery / Cloud Storage)
5. Explore & Visualize (DataStudio)

# Message oriented architecture
Used for data ingestion
Late, incorrect, ...
Real time Iot 

# Cloud dataflow
Reçoit un job et exécute le job
Des templates de jobs existent pour PubSub to BigQuery, ...

# DataStudio
Dimensions (green chips): Categories
Metrics (Blue chips): Measure dimension values
Calculated fields: Used to create metrics or dimensions (Regexp to extract landing page URL)
