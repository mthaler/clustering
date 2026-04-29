# Install ggplot2 if not already installed
# install.packages("ggplot2")

# load the ggplot 2 library
library(ggplot2)

# calculate kmeans of the dataset
kmeans_result <- kmeans(faithful, centers = 2, nstart = 10)

# Add cluster assignments to the original data
faithful$Cluster <- as.factor(kmeans_result$cluster)

# Create a data frame for cluster centers
centers <- as.data.frame(kmeans_result$centers)
centers$Cluster <- as.factor(1:2)

# Plot the data using ggplot2
ggplot(iris, aes(x = Sepal.Length, y = Sepal.Width, color = Cluster)) +
  geom_point(size = 2, alpha = 0.7) +
  geom_point(
    data = centers,
    aes(x = Sepal.Length, y = Sepal.Width),
    color = "red",
    shape = 8,
    size = 5,
    stroke = 2
  ) +
  labs(
    title = "Visualizing Clusters in R using Old Faithful dataset",
    x = "Sepal Length (cm)",
    y = "Sepal Width (cm)",
    color = "Cluster"
  ) +
  theme_bw() +
  guides(color = guide_legend(override.aes = list(size = 4)))