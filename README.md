# online-diff
Investigating the similarity of the texts on the uploaded batch of text files

Go web-service (Minio, SQLite), Python NLP (Glove, gensim, nltk, diff-match-patch) – web-service allows investigating
the similarity of the texts on the uploaded batch of text files:

○ Compute matrix of pairwise semantic/syntactic similarities using soft-cosine/cosine similarity algorithms;

○ Get the difference between two uploaded early documents using Myer’s diff algorithm;
