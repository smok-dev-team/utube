CREATE TABLE `youtube_channel` (
  `id` varchar(128) NOT NULL,
  `title` varchar(512) DEFAULT NULL,
  `description` varchar(2048) DEFAULT NULL,
  `thumbnails` varchar(1024) DEFAULT NULL,
  `published_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `youtube_channel_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `youtube_video` (
  `id` varchar(128) NOT NULL,
  `channel_id` varchar(128) DEFAULT NULL,
  `playlist_id` varchar(128) DEFAULT NULL,
  `title` varchar(512) DEFAULT NULL,
  `description` varchar(2048) DEFAULT NULL,
  `thumbnails` varchar(1024) DEFAULT NULL,
  `published_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `youtube_video_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

