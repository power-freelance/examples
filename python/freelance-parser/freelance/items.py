# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


PRICE_PERIOD_HOURLY = 'hourly'
PRICE_PERIOD_PROJECT = 'project'

class TaskItem(scrapy.Item):
    url = scrapy.Field()
    title = scrapy.Field()
    views = scrapy.Field(serializer=int)
    responses = scrapy.Field(serializer=int)
    is_safe_deal = scrapy.Field(serializer=bool)
    tags = scrapy.Field(serializer=list)
    price_negotiated = scrapy.Field(serializer=bool)
    price_amount = scrapy.Field(serializer=int)
    price_period = scrapy.Field()
