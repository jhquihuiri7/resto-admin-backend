cluster = AgglomerativeClustering(n_clusters=4, linkage="ward")
cluster.fit(df)

pca = PCA(n_components=2)
X_pca = pca.fit_transform(df)

plt.figure(figsize=(10,10))
plt.scatter(X_pca[:,0], X_pca[:,1], c=cluster.labels_, cmap="viridis")
plt.title("Clustering Aglomerativo con PCA (2D)")
plt.xlabel("Componente Principal 1")
plt.ylabel("Componente Principal 2")
plt.colorbar(label="Cluster")
plt.show()