from django.urls import path

from . import views

urlpatterns = [
    path('getEntries/', views.getEntries, name='getEntries'),
]